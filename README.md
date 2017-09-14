# ifog
Go client for iCloud rest APIs

Only device info retrieval is implemented.
I plan to add photo API later (currently the icloud API has been deactivated by Apple ...)

Example :
```go

    session := ifog.NewSession()
	err := session.Login(login.RequestBody{"myAppleId", "myPassword"})
	if err != nil {
		panic(err)
	}
	error := session.PopulateDevices()
	if error != nil {
		panic(error)
	}
	myDevice := session.Devices[0]
	fmt.Println("My " + myDevice.DeviceClass + " is located at : " + string(myDevice.Location.Latitude) + ";" + string(myDevice.Location.Longitude))
```


Printing :

```shell
    My iPhone is located at : 43.107062, 0.723592
```
