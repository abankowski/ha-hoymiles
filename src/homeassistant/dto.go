package homeassistant

type Device struct {
	ConfigurationUrl *string `json:"configuration_url,omitempty"`
	Connections      *string `json:"connections,omitempty"`
	HwVersion        *string `json:"hw_version,omitempty"`
	Identifiers      *string `json:"identifiers,omitempty"`
	Manufacturer     *string `json:"manufacturer,omitempty"`
	Model            *string `json:"model,omitempty"`
	Name             *string `json:"name,omitempty"`
	SuggestedArea    *string `json:"suggested_area,omitempty"`
	SwVersion        *string `json:"sw_version,omitempty"`
	ViaDevice        *string `json:"via_device,omitempty"`
}

type Availability struct {
	Topic               *string `json:"name"`
	PayloadAvailable    *string `json:"payload_available,omitempty"`
	PayloadNotAvailable *string `json:"payload_not_available,omitempty"`
	ValueTemplate       *string `json:"value_template,omitempty"`
}

type Sensor struct {
	Availability           *[]Availability `json:"availability,omitempty"`
	Name                   *string         `json:"name"`
	NativeValue            *string         `json:"native_value"`
	DeviceClass            *string         `json:"device_class"`
	Icon                   *string         `json:"icon"`
	UnitOfMeasurement      *string         `json:"unit_of_measurement"`
	Device                 *Device
	AvailabilityTopic      *string `json:"availability_topic,omitempty"`
	AvailabilityTemplate   *string `json:"availability_template,omitempty"`
	AvailabilityMode       *string `json:"availability_mode,omitempty"`
	EnabledByDefault       *bool   `json:"enabled_by_default,omitempty"`
	Encoding               *string `json:"encoding,omitempty"`
	EntityCategory         *string `json:"entity_category,omitempty"`
	ExpireAfter            *int    `json:"expire_after,omitempty"`
	ForceUpdate            *bool   `json:"force_update,omitempty"`
	JsonAttributesTemplate *string `json:"json_attributes_template,omitempty"`
	JsonAttributesTopic    *string `json:"json_attributes_topic,omitempty"`
	LastResetValueTemplate *string `json:"last_reset_value_template,omitempty"`
	ObjectId               *string `json:"object_id,omitempty"`
	PayloadAvailable       *string `json:"payload_available,omitempty"`
	PayloadNotAvailable    *string `json:"payload_not_available,omitempty"`
	Qos                    *int    `json:"qos,omitempty"`
	StateClass             *string `json:"state_class,omitempty"`
	StateTopic             *string `json:"state_topic,omitempty"`
	UniqueId               *string `json:"unique_id,omitempty"`
	ValueTemplate          *string `json:"value_template,omitempty"`
	Rssi                   *string `json:"RSSI,omitempty"`
}

type Dtu struct {
	BinarySensor struct {
		Connect Sensor `json:"connect"`
		Total   Sensor `json:"total"`
	} `json:"binary_sensor"`
}

type MicroInverter struct {
	BinarySensor struct {
		Connect Sensor `json:"connect"`
		Total   Sensor `json:"total"`
	} `json:"binary_sensor"`

	Sensor struct {
		AlarmCode   Sensor `json:"alarm_code"`
		AlarmString Sensor `json:"alarm_string"`
		Total       Sensor `json:"sensor"`
	} `json:"sensor"`
}

type SolarPlant struct {
	Sensor struct {
		Capacitor Sensor `json:"capacitor"`

		CapacitorKW              Sensor `json:"capacitor_kW"`
		Co2EmissionReduction     Sensor `json:"co2_emission_reduction"`
		IsBalance                Sensor `json:"is_balance"`
		IsReflux                 Sensor `json:"is_reflux"`
		LastDataTime             Sensor `json:"last_data_time"`
		MonthEq                  Sensor `json:"month_eq"`
		PlantTree                Sensor `json:"plant_tree"`
		RealPower                Sensor `json:"real_power"`
		RealPowerKw              Sensor `json:"real_power_kw"`
		RealPowerMeasurement     Sensor `json:"real_power_measurement"`
		RealPowerTotalIncreasing Sensor `json:"real_power_total_increasing"`
		RefluxStationData        Sensor `json:"reflux_station_data"`
		TodayEq                  Sensor `json:"today_eq"`
		TodayEqWh                Sensor `json:"today_eq_Wh"`
		YearEq                   Sensor `json:"year_eq"`
		TotalEq                  Sensor `json:"total_eq"`
		PowerRatio               Sensor `json:"power_ratio"`
		Total                    Sensor `json:"total"`
	} `json:"sensor"`
}

type MicroInverterPort struct {
	Capacitor                Sensor `json:"capacitor"`
	CapacitorKW              Sensor `json:"capacitor_kW"`
	Co2EmissionReduction     Sensor `json:"co2_emission_reduction"`
	IsBalance                Sensor `json:"is_balance"`
	IsReflux                 Sensor `json:"is_reflux"`
	LastDataTime             Sensor `json:"last_data_time"`
	MonthEq                  Sensor `json:"month_eq"`
	PlantTree                Sensor `json:"plant_tree"`
	RealPower                Sensor `json:"real_power"`
	RealPowerKw              Sensor `json:"real_power_kw"`
	RealPowerMeasurement     Sensor `json:"real_power_measurement"`
	RealPowerTotalIncreasing Sensor `json:"real_power_total_increasing"`
	RefluxStationData        Sensor `json:"reflux_station_data"`
	TodayEq                  Sensor `json:"today_eq"`
	TodayEqWh                Sensor `json:"today_eq_Wh"`
	TotalEq                  Sensor `json:"total_eq"`
	PowerRatio               Sensor `json:"power_ratio"`
	Total                    Sensor `json:"total"`
}
