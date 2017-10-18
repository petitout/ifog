[![Build Status](https://travis-ci.org/petitout/ifog.svg?branch=master)](https://travis-ci.org/petitout/ifog)
[![Coverage Status](https://coveralls.io/repos/github/petitout/ifog/badge.svg?branch=master)](https://coveralls.io/github/petitout/ifog?branch=master)

# ifog
Go client for iCloud rest APIs

Only device info retrieval is implemented.
I plan to add photo API later (currently the icloud API has been deactivated by Apple ...)

Example :
```go
	session := ifog.NewSession()
	var err = session.Login(login.RequestBody{"myAppleId", "myPassword"})
	if err != nil {
		panic(err)
	}
	err = session.PopulateDevices()
	if err != nil {
		panic(err)
	}
	myDevice := session.Devices[0]
	fmt.Println("My " + myDevice.DeviceClass + " is located at : " + myDevice.Location.String())
```


Printing :

```shell
    My iPhone is located at : 43.107062, 0.723592
```
