# sentry-client
Consuming Specific APIs from Sentry

Current progress : 


```
[
{
	"eventID": "FBA551EA3CAF44B8BE37695873F7868B",
	"packages": {},
	"dist": "1",
	"tags": [
		{
			"value": "151f43cd37f13264841ad75d7e5c0fd74a5a340b",
			"key": "app.device"
		}
	],
	"contexts": {
		"device": {
			"model_id": "D201AP",
			"family": "iOS",
			"arch": "arm64",
			"storage_size": 63989469184,
			"free_memory": 321667072,
			"memory_size": 2086879232,
			"boot_time": "2019-06-04T11:25:47Z",
			"timezone": "MESZ",
			"model": "iPhone10,4",
			"usable_memory": 1235642944,
			"type": "device"
		},
		"app": {
			"app_identifier": "com.linux.myapp",
			"app_name": "MyApp",
			"device_app_hash": "151f43cd37f1216e65t5x73264841ad75d7e5c0fd74a5a340b",
			"app_id": "0ED11752-92A7-3F81-AC7E-123434D7774",
			"build_type": "test",
			"app_start_time": "2019-06-06T12:25:17Z",
			"app_version": "2.6.0",
			"type": "app",
			"app_build": "1"
		},
		"os": {
			"version": "12.3.1",
			"rooted": false,
			"name": "iOS",
			"kernel_version": "Darwin Kernel",
			"type": "os",
			"build": "16F203"
		}
	},
	"dateReceived": "2019-06-06T14:16:55Z",
	"dateCreated": "2019-06-06T14:16:55Z",
	"fingerprints": [
		"ad3182375fcasas89fda014b6873d147e"
	],
	"metadata": {
		"title": "title here"
	},
	"groupID": "557",
	"platform": "cocoa",
	"errors": [],
	"user": {
		"username": "finformauto",
		"id": "XXX-XXXXX"
	},
	"context": {
		"Duration": "00:26"
	},
	"entries": [
		{
			"type": "message",
			"data": {
				"message": "Ident finished"
			}
		},
		{
			"type": "breadcrumbs",
			"data": {
				"values": [
					{
						"category": "Docket",
						"level": "info",
						"event_id": null,
						"timestamp": "2019-06-06T14:16:46Z",
						"data": null,
						"message": "<< Received",
						"type": "default"
					}
				]
			}
		}
	],
	"message": "Ident finished",
	"sdk": {
		"version": "4.1.0",
		"name": "sentry-cocoa",
		"upstream": {
			"url": null,
			"isNewer": false,
			"name": "sentry-cocoa"
		}
	},
	"type": "default",
	"id": "30307",
	"size": 27236
}
]
```


```
type event struct {
	EventID string `json:"eventID"`
	Packages packages`json:"packages"`
	Dist string`json:"dist"`
	Tags []tags`json:"tags"`
	Contexts contexts`json:"contexts"`
	DateReceived string`json:"dateReceived"`
	DateCreated string`json:"dateReceived"`
	Fingerprints []string`json:"fingerprints"`
	Metadata metadata`json:"metadata"`
	GroupID string`json:"groupID"`
	Platform string`json:"platform"`
	Errors []string`json:"errors"`
	User user`json:"user"`
	Context context`json:"context"`
	Entries []entry`json:"entries"`

}

type packages struct {

}

type tags struct {
	Value string`json:"value"`
	Key string`json:"key"`
}

type contexts struct{
	Device device`json:"device"`
	App app`json:"app"`
	OperatingSystem operatingSystem`json:"os"`

}

type device struct {
	ModelID string`json:"model_id"`
	Family string`json:"family"`
	Arch string`json:"arch"`
	StorageSize int`json:"storage_size"`
	FreeMemory int`json:"free_memory"`
	MemorySize int`json:"memory_size"`
	BootTime string`json:"boot_time"`
	Timezone string`json:"timezone"`
	Model string`json:"model"`
	UsableMemory int`json:"usable_memory"`
	Type string`json:"type"`
}

type app struct {
	AppIdentifier string`json:"app_identifier"`
	AppName string`json:"app_name"`
	DeviceAppHash string`json:"device_app_hash"`
	AppId string`json:"app_id"`
	BuildType string`json:"build_type"`
	AppStartTime string`json:"app_start_time"`
	AppVersion string`json:"app_version"`
	Type string`json:"type"`
	AppBuild int`json:"app_build"`
}

type operatingSystem struct {
	Version string `json:"version"`
	Rooted bool `json:"rooted"`
	Name string `json:"name"`
	KernelVersion string `json:"kernel_version"`
	Type string `json:"type"`
	Build string `json:"build"`
}

type metadata struct {
	Title string`json:"title"`
}

type user struct {
	Shortname string`json:"username"`
	InternalToken string`json:"id"`
}

type context struct {
	Duration string`json:"Duration"`
}

type entry struct {
	Type string`json:"type"`
	Data data`json:"data"`
}

type data struct {
	Message string`json:"message"`
}
```
