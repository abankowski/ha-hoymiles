// Package hoymiles provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.11.0 DO NOT EDIT.
package hoymiles

const (
	CookieAuthScopes = "cookieAuth.Scopes"
)

// Station Find Request
type AuthLoginRequest struct {
	ERRORBACK      bool `json:"ERROR_BACK"`
	WAITINGPROMISE bool `json:"WAITING_PROMISE"`
	Body           struct {
		Password string `json:"password"`
		UserName string `json:"user_name"`
	} `json:"body"`
}

// Auth Login response with token
type AuthLoginResponse struct {
	Data *struct {
		Token string `json:"token"`
	} `json:"data,omitempty"`
	Message      string  `json:"message"`
	Status       string  `json:"status"`
	SystemNotice *string `json:"systemNotice,omitempty"`
}

// Station Find Response envelope
type DataCountStationRealDataResponse struct {
	Data *struct {
		Capacitor            *string `json:"capacitor,omitempty"`
		Co2EmissionReduction *string `json:"co2_emission_reduction,omitempty"`
		DataTime             *string `json:"data_time,omitempty"`
		IsBalance            *int    `json:"is_balance,omitempty"`
		IsReflux             *int    `json:"is_reflux,omitempty"`
		LastDataTime         *string `json:"last_data_time,omitempty"`
		MonthEq              *string `json:"month_eq,omitempty"`
		PlantTree            *string `json:"plant_tree,omitempty"`
		RealPower            *string `json:"real_power,omitempty"`
		TodayEq              *string `json:"today_eq,omitempty"`
		TotalEq              *string `json:"total_eq,omitempty"`
		YearEq               *string `json:"year_eq,omitempty"`
	} `json:"data,omitempty"`
	Message      *string `json:"message,omitempty"`
	Status       *string `json:"status,omitempty"`
	SystemNotice *string `json:"systemNotice,omitempty"`
}

// DataFindDetailsRequest defines model for DataFindDetailsRequest.
type DataFindDetailsRequest struct {
	WAITINGPROMISE bool `json:"WAITING_PROMISE"`
	Body           struct {
		MiId     int    `json:"mi_id"`
		MiSn     string `json:"mi_sn"`
		Port     int    `json:"port"`
		Sid      int    `json:"sid"`
		Time     string `json:"time"`
		WarnCode int    `json:"warn_code"`
	} `json:"body"`
}

// DataFindDetailsResponse defines model for DataFindDetailsResponse.
type DataFindDetailsResponse struct {
	Data *struct {
		Dbg      *bool   `json:"dbg,omitempty"`
		MiSn     *string `json:"mi_sn,omitempty"`
		NetState *int    `json:"net_state,omitempty"`
		WarnList *[]struct {
			EndAt   *string `json:"end_at,omitempty"`
			ErrCode *int    `json:"err_code,omitempty"`
			StartAt *string `json:"start_at,omitempty"`
			Wd1     *string `json:"wd1,omitempty"`
			Wd2     *string `json:"wd2,omitempty"`
			Wdd1    *string `json:"wdd1,omitempty"`
			Wdd2    *string `json:"wdd2,omitempty"`
		} `json:"warn_list,omitempty"`
	} `json:"data,omitempty"`
	Message      *string `json:"message,omitempty"`
	Status       *string `json:"status,omitempty"`
	SystemNotice *string `json:"systemNotice,omitempty"`
}

// DeviceWarnData defines model for DeviceWarnData.
type DeviceWarnData struct {
	Rw      *string `json:"_rw,omitempty"`
	Connect bool    `json:"connect"`
	Warn    bool    `json:"warn"`
}

// Get authenticated user details
type GetUserMeRequest struct {
	ERRORBACK      bool `json:"ERROR_BACK"`
	WAITINGPROMISE bool `json:"WAITING_PROMISE"`
}

// Autenthicated user response payload
type GetUserMeResponse struct {
	Data *struct {
		CityId   *int    `json:"city_id,omitempty"`
		CreateAt *string `json:"create_at,omitempty"`
		CreateBy *int    `json:"create_by,omitempty"`
		Email    *string `json:"email,omitempty"`
		Gid      *int    `json:"gid,omitempty"`
		Id       *int    `json:"id,omitempty"`
		Identity *int    `json:"identity,omitempty"`
		Name     *string `json:"name,omitempty"`
		Phone    *string `json:"phone,omitempty"`
		RoleIds  *[]int  `json:"role_ids,omitempty"`
		Src      *int    `json:"src,omitempty"`
		Status   *int    `json:"status,omitempty"`
		Tid      *int    `json:"tid,omitempty"`
		Type     *int    `json:"type,omitempty"`
		UpdateAt *string `json:"update_at,omitempty"`
		UpdateBy *int    `json:"update_by,omitempty"`
		UserName *string `json:"user_name,omitempty"`
	} `json:"data,omitempty"`
	Message      *string `json:"message,omitempty"`
	Status       *string `json:"status,omitempty"`
	SystemNotice *string `json:"systemNotice,omitempty"`
}

// ParentCity defines model for ParentCity.
type ParentCity struct {
	CityName     *string `json:"city_name,omitempty"`
	Code         *string `json:"code,omitempty"`
	CountryCode  *string `json:"country_code,omitempty"`
	Id           *int    `json:"id,omitempty"`
	Level        *int    `json:"level,omitempty"`
	Pid          *int    `json:"pid,omitempty"`
	WeatherOfCid *int    `json:"weather_of_cid,omitempty"`
}

// Station Data Request
type StationDataRequest struct {
	WAITINGPROMISE bool `json:"WAITING_PROMISE"`
	Body           struct {
		Sid int `json:"sid"`
	} `json:"body"`
}

// StationDesc config subobject
type StationDescConfig struct {
	BillingEvery      *int    `json:"billing_every,omitempty"`
	BillingStart      *string `json:"billing_start,omitempty"`
	BillingType       *int    `json:"billing_type,omitempty"`
	ClientType        *int    `json:"client_type,omitempty"`
	ModuleMaxPower    *int    `json:"module_max_power,omitempty"`
	OwnerIsModifyDev  *int    `json:"owner_is_modify_dev,omitempty"`
	OwnerIsShowLayout *int    `json:"owner_is_show_layout,omitempty"`
	PowerLimit        *string `json:"power_limit,omitempty"`
	PowerLimitPf      *string `json:"power_limit_pf,omitempty"`
	PowerLimitRe      *string `json:"power_limit_re,omitempty"`
	SunSpecNum        *int    `json:"sun_spec_num,omitempty"`
}

// Station Find Response envelope
type StationFindResponse struct {
	Data *struct {
		Address      *string `json:"address,omitempty"`
		BmsCapacitor *string `json:"bms_capacitor,omitempty"`
		Capacitor    *string `json:"capacitor,omitempty"`
		CityCode     *string `json:"city_code,omitempty"`
		CityId       *int    `json:"city_id,omitempty"`
		Classify     *int    `json:"classify,omitempty"`

		// StationDesc config subobject
		Config           *StationDescConfig `json:"config,omitempty"`
		CreateAt         *string            `json:"create_at,omitempty"`
		CreateBy         *int               `json:"create_by,omitempty"`
		ElectricityPrice *int               `json:"electricity_price,omitempty"`
		Gid              *int               `json:"gid,omitempty"`
		Group            *StationGroup      `json:"group,omitempty"`
		Id               int                `json:"id"`
		InPrice          *float32           `json:"in_price,omitempty"`
		IsBalance        *int               `json:"is_balance,omitempty"`
		IsReflux         *int               `json:"is_reflux,omitempty"`
		IsStars          *int               `json:"is_stars,omitempty"`
		Latitude         *string            `json:"latitude,omitempty"`
		LayoutStep       *int               `json:"layout_step,omitempty"`
		LocalTime        *string            `json:"local_time,omitempty"`
		Longitude        *string            `json:"longitude,omitempty"`
		MeterLocation    *int               `json:"meter_location,omitempty"`
		MoneyData        *struct {
			Code string `json:"code"`
			Unit string `json:"unit"`
		} `json:"money_data,omitempty"`
		MoneyUnit    *string          `json:"money_unit,omitempty"`
		Name         *string          `json:"name,omitempty"`
		OwnerList    *[]StationOwner  `json:"owner_list,omitempty"`
		ParentCity   *[]ParentCity    `json:"parent_city,omitempty"`
		PicPath      *string          `json:"pic_path,omitempty"`
		Remarks      *string          `json:"remarks,omitempty"`
		Status       *int             `json:"status,omitempty"`
		Timezone     *Timezone        `json:"timezone,omitempty"`
		Type         *int             `json:"type,omitempty"`
		TzId         *int             `json:"tz_id,omitempty"`
		TzName       *string          `json:"tz_name,omitempty"`
		Usd          *string          `json:"usd,omitempty"`
		WarnData     *StationWarnData `json:"warn_data,omitempty"`
		WeatherOfCid *int             `json:"weather_of_cid,omitempty"`
	} `json:"data,omitempty"`
	Message      *string `json:"message,omitempty"`
	Status       *string `json:"status,omitempty"`
	SystemNotice *string `json:"systemNotice,omitempty"`
}

// StationGroup defines model for StationGroup.
type StationGroup struct {
	Area    *string `json:"area,omitempty"`
	Contact *string `json:"contact,omitempty"`
	Icon    *string `json:"icon,omitempty"`
	Id      *int    `json:"id,omitempty"`
	Name    *string `json:"name,omitempty"`
	Phone   *string `json:"phone,omitempty"`
	Pid     *int    `json:"pid,omitempty"`
	Tid     *int    `json:"tid,omitempty"`
}

// StationOwner defines model for StationOwner.
type StationOwner struct {
	Name     *string `json:"name,omitempty"`
	Uid      *int    `json:"uid,omitempty"`
	UserName *string `json:"user_name,omitempty"`
}

// Station Request
type StationRequest struct {
	WAITINGPROMISE bool `json:"WAITING_PROMISE"`
	Body           struct {
		Id int `json:"id"`
	} `json:"body"`
}

// StationSelectDeviceAllResponse
type StationSelectDeviceAllResponse struct {
	Data *[]struct {
		Dtu *struct {
			HardVer    *string         `json:"hard_ver,omitempty"`
			Id         *int            `json:"id,omitempty"`
			ModelNo    *string         `json:"model_no,omitempty"`
			ReplaceNum *int            `json:"replace_num,omitempty"`
			Sn         *string         `json:"sn,omitempty"`
			SoftVer    *string         `json:"soft_ver,omitempty"`
			Version    *int            `json:"version,omitempty"`
			WarnData   *DeviceWarnData `json:"warn_data,omitempty"`
		} `json:"dtu,omitempty"`
		RepeaterList *[]struct {
			Micros *[]StationSelectDeviceMicro `json:"micros,omitempty"`
		} `json:"repeater_list,omitempty"`
	} `json:"data,omitempty"`
	Message      *string `json:"message,omitempty"`
	Status       *string `json:"status,omitempty"`
	SystemNotice *string `json:"systemNotice,omitempty"`
}

// StationSelectDeviceMicro defines model for StationSelectDeviceMicro.
type StationSelectDeviceMicro struct {
	DevType     *int    `json:"dev_type,omitempty"`
	DtuSn       *string `json:"dtu_sn,omitempty"`
	GridId      *int    `json:"grid_id,omitempty"`
	GridName    *string `json:"grid_name,omitempty"`
	GridVersion *string `json:"grid_version,omitempty"`
	HardVer     *string `json:"hard_ver,omitempty"`
	Id          *int    `json:"id,omitempty"`
	InitHardNo  *string `json:"init_hard_no,omitempty"`
	PortArray   *[]int  `json:"port_array,omitempty"`
	ReplaceNum  *int    `json:"replace_num,omitempty"`
	Series      *string `json:"series,omitempty"`
	SeriesName  *string `json:"series_name,omitempty"`
	Sn          *string `json:"sn,omitempty"`
	SoftVer     *string `json:"soft_ver,omitempty"`
	Version     *int    `json:"version,omitempty"`
	WarnData    *struct {
		DtuLink   *bool   `json:"_dtu_link,omitempty"`
		LTime     *int    `json:"_l_time,omitempty"`
		RawWarns  *string `json:"_raw_warns,omitempty"`
		Connect   *bool   `json:"connect,omitempty"`
		TrackTime *string `json:"track_time,omitempty"`
		Warn      *bool   `json:"warn,omitempty"`
	} `json:"warn_data,omitempty"`
}

// StationSelectDeviceOfTree defines model for StationSelectDeviceOfTree.
type StationSelectDeviceOfTree struct {
	Children   *[]StationSelectDeviceOfTree `json:"children,omitempty"`
	DtuSn      *string                      `json:"dtu_sn,omitempty"`
	HardVer    *string                      `json:"hard_ver,omitempty"`
	Id         *int                         `json:"id,omitempty"`
	ModelNo    *string                      `json:"model_no,omitempty"`
	ReplaceNum *int                         `json:"replace_num,omitempty"`
	Sn         *string                      `json:"sn,omitempty"`
	SoftVer    *string                      `json:"soft_ver,omitempty"`
	Type       *int                         `json:"type,omitempty"`
	Version    *int                         `json:"version,omitempty"`
	WarnData   *DeviceWarnData              `json:"warn_data,omitempty"`
}

// Station Select Device Tree Response envelope
type StationSelectDeviceOfTreeResponse struct {
	Data         []StationSelectDeviceOfTree `json:"data"`
	Message      string                      `json:"message"`
	Status       string                      `json:"status"`
	SystemNotice *string                     `json:"systemNotice,omitempty"`
}

// StationWarnData defines model for StationWarnData.
type StationWarnData struct {
	GWarn    *bool `json:"g_warn,omitempty"`
	L3Warn   *bool `json:"l3_warn,omitempty"`
	MeWarn   *bool `json:"me_warn,omitempty"`
	SUid     *bool `json:"s_uid,omitempty"`
	SUoff    *bool `json:"s_uoff,omitempty"`
	SUstable *bool `json:"s_ustable,omitempty"`
}

// Timezone defines model for Timezone.
type Timezone struct {
	DisName *string `json:"dis_name,omitempty"`
	Id      *int    `json:"id,omitempty"`
	Name    *string `json:"name,omitempty"`
	Offset  *int    `json:"offset,omitempty"`
	TzName  *string `json:"tz_name,omitempty"`
}

// PostIamAuthLoginJSONBody defines parameters for PostIamAuthLogin.
type PostIamAuthLoginJSONBody = AuthLoginRequest

// PostIamUserMeJSONBody defines parameters for PostIamUserMe.
type PostIamUserMeJSONBody = GetUserMeRequest

// PostPvmDataDataCountStationRealDataJSONBody defines parameters for PostPvmDataDataCountStationRealData.
type PostPvmDataDataCountStationRealDataJSONBody = StationDataRequest

// PostPvmDataDataFindDetailsJSONBody defines parameters for PostPvmDataDataFindDetails.
type PostPvmDataDataFindDetailsJSONBody = DataFindDetailsRequest

// PostPvmStationFindJSONBody defines parameters for PostPvmStationFind.
type PostPvmStationFindJSONBody = StationRequest

// PostPvmStationSelectDeviceAllJSONBody defines parameters for PostPvmStationSelectDeviceAll.
type PostPvmStationSelectDeviceAllJSONBody = StationRequest

// PostPvmStationSelectDeviceOfTreeJSONBody defines parameters for PostPvmStationSelectDeviceOfTree.
type PostPvmStationSelectDeviceOfTreeJSONBody = StationRequest

// PostIamAuthLoginJSONRequestBody defines body for PostIamAuthLogin for application/json ContentType.
type PostIamAuthLoginJSONRequestBody = PostIamAuthLoginJSONBody

// PostIamUserMeJSONRequestBody defines body for PostIamUserMe for application/json ContentType.
type PostIamUserMeJSONRequestBody = PostIamUserMeJSONBody

// PostPvmDataDataCountStationRealDataJSONRequestBody defines body for PostPvmDataDataCountStationRealData for application/json ContentType.
type PostPvmDataDataCountStationRealDataJSONRequestBody = PostPvmDataDataCountStationRealDataJSONBody

// PostPvmDataDataFindDetailsJSONRequestBody defines body for PostPvmDataDataFindDetails for application/json ContentType.
type PostPvmDataDataFindDetailsJSONRequestBody = PostPvmDataDataFindDetailsJSONBody

// PostPvmStationFindJSONRequestBody defines body for PostPvmStationFind for application/json ContentType.
type PostPvmStationFindJSONRequestBody = PostPvmStationFindJSONBody

// PostPvmStationSelectDeviceAllJSONRequestBody defines body for PostPvmStationSelectDeviceAll for application/json ContentType.
type PostPvmStationSelectDeviceAllJSONRequestBody = PostPvmStationSelectDeviceAllJSONBody

// PostPvmStationSelectDeviceOfTreeJSONRequestBody defines body for PostPvmStationSelectDeviceOfTree for application/json ContentType.
type PostPvmStationSelectDeviceOfTreeJSONRequestBody = PostPvmStationSelectDeviceOfTreeJSONBody
