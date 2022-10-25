package main

import (
	"context"
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/procyon-projects/chrono"
	"github.com/spf13/viper"
	"ha-hoymiles/hoymiles"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// TODO: Move main content to package
func main() {
	if viper.BindEnv("MQTT_HOST") != nil || viper.BindEnv("MQTT_USER") != nil || viper.BindEnv("MQTT_PASSWORD") != nil {
		log.Panic("Some errors in binding env variables")
	}

	taskScheduler := chrono.NewDefaultTaskScheduler()

	mqttOpts := mqtt.NewClientOptions()

	mqttOpts.AddBroker(viper.GetString("mqtt_host"))
	mqttOpts.SetUsername(viper.GetString("mqtt_user"))
	mqttOpts.SetPassword(viper.GetString("mqtt_password"))
	mqttOpts.SetAutoReconnect(true)
	mqttClient := mqtt.NewClient(mqttOpts)
	mqttClient.Connect()

	task, err := taskScheduler.ScheduleAtFixedRate(func(ctx context.Context) {
		if mqttClient.IsConnectionOpen() {
			updateAll(ctx, mqttClient)
		} else {
			log.Fatalf("Cannot execute update loop when mqtt is not connected, aborting")
		}
	}, 5*time.Minute) // TODO: make this configurable

	if err != nil {
		log.Panic(err)
	}

	sigc := make(chan os.Signal, 1) // Magic value since this will be the last message in this channel

	signal.Notify(sigc,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT)

	<-sigc
	task.Cancel()
	mqttClient.Disconnect(1000) //TODO: pick reasonable value and place it in a constant instead of magic value
	shutdownChannel := taskScheduler.Shutdown()
	<-shutdownChannel
}

// TODO: move to package
func updateAll(ctx context.Context, mqttClient mqtt.Client) {

	// TODO: take this from configuration
	station := viper.GetInt("hoymiles_station_id")
	c, err := hoymiles.NewClient(
		viper.GetString("hoymiles_username"),
		viper.GetString("hoymiles_password"),
		ctx,
		true)

	if err != nil {
		log.Fatal("cannot get client with auth", err)
	}

	tree, err := c.GetDeviceTree(station)

	for _, dtu := range tree.Data {

		fmt.Printf("DTU %s %s %s %s\n", *dtu.Sn, *dtu.ModelNo, *dtu.HardVer, *dtu.SoftVer)

		for _, mi := range *dtu.Children {
			fmt.Printf("  %s %s %s %s\n", *mi.Sn, *mi.ModelNo, *dtu.HardVer, *mi.SoftVer)

			t := time.Now().Format("2006-01-02 15:04")

			dd, _ := c.GetDataDetails(station, *mi.Id, *mi.Sn, t)

			for _, w := range *dd.Data.WarnList {
				if w.ErrCode != nil {
					fmt.Printf("    error=%d [%s] %s %s\n", *w.ErrCode, hoymiles.StatusCodes[*w.ErrCode], *w.Wdd1, *w.Wd1)
				}
			}
		}
	}

	//
	//fmt.Println(erx)
	//
	//u, err := c.GetUserMe()
	//
	//if err != nil {
	//
	//	log.Fatal("cannot get user me ", err)
	//}
	//fmt.Printf("I am %s %s\n", *u.Data.Name, *u.Data.Email)
	//
	//_, err = c.GetRealData(station)
	//
	//if err != nil {
	//	log.Fatal("cannot get real data", err)
	//}
	//
	//_, err = c.GetStation(station)
	//
	//if err != nil {
	//	log.Fatal("cannot get station", err)
	//}
	//
	//_, err = c.GetDeviceTree(station)
	//
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//_, err = c.GetDeviceAll(station)
	//
	//if err != nil {
	//	log.Fatal("cannot get station", err)
	//}

}
