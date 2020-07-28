# Philips Hue Sample

Refs: <https://developers.meethue.com/develop/get-started-2/>

```sh
$ cp hue-config.json.sample hue-config.json
$ vi hue-config.json # edit this
$ go run . # up server

---

$ curl localhost:8080/dance
```

## Hue API の使い方

1. IPアドレスを調べる

ネイティブアプリHue

設定 => Hueブリッジ => i => IPアドレス

2. Hueブリッジにユーザー登録する

```sh
$ curl -s -X POST -d "$(jo devicetype=ore-no-iot)" http://<-ip-address->/api | jq .
[
  {
    "error": {
      "type": 101,
      "address": "",
      "description": "link button not pressed"
    }
  }
]
```

Hueブリッジ本体中心のボタンを押す

```sh
$ curl -s -X POST -d "$(jo devicetype=ore-no-iot)" http://<-ip-address->/api | jq .
[
  {
    "success": {
      "username": "<-user-name->"
    }
  }
]
```

3. 登録されているライトを一覧

```sh
$ curl -s http://<-ip-address->/api/<-user-name->/lights | jq .
{
  "1": {
    "state": {
      "on": true,
      "bri": 254,
      "ct": 366,
      "alert": "select",
      "colormode": "ct",
      "mode": "homeautomation",
      "reachable": true
    },
    "swupdate": {
      "state": "noupdates",
      "lastinstall": "2020-07-28T06:40:42"
    },
    "type": "Color temperature light",
    "name": "Hue ambiance lamp 1",
    "modelid": "LTW001",
    "manufacturername": "Signify Netherlands B.V.",
    "productname": "Hue ambiance lamp",
    "capabilities": {
      "certified": true,
      "control": {
        "mindimlevel": 1000,
        "maxlumen": 806,
        "ct": {
          "min": 153,
          "max": 454
        }
      },
      "streaming": {
        "renderer": false,
        "proxy": false
      }
    },
    "config": {
      "archetype": "sultanbulb",
      "function": "functional",
      "direction": "omnidirectional",
      "startup": {
        "mode": "safety",
        "configured": true
      }
    },
    "uniqueid": "00:00:00:00:00:00:00:00-00",
    "swversion": "0.000.0.00000"
  }
}
```

4. 登録されているライトをIDで取得

```sh
$ curl -s http://<-ip-address->/api/<-user-name->/lights/1 | jq .
{
  "state": {
    "on": true,
    "bri": 254,
    "ct": 366,
    "alert": "select",
    "colormode": "ct",
    "mode": "homeautomation",
    "reachable": true
  },
  "swupdate": {
    "state": "noupdates",
    "lastinstall": "2020-07-28T06:40:42"
  },
  "type": "Color temperature light",
  "name": "Hue ambiance lamp 1",
  "modelid": "LTW001",
  "manufacturername": "Signify Netherlands B.V.",
  "productname": "Hue ambiance lamp",
  "capabilities": {
    "certified": true,
    "control": {
      "mindimlevel": 1000,
      "maxlumen": 806,
      "ct": {
        "min": 153,
        "max": 454
      }
    },
    "streaming": {
      "renderer": false,
      "proxy": false
    }
  },
  "config": {
    "archetype": "sultanbulb",
    "function": "functional",
    "direction": "omnidirectional",
    "startup": {
      "mode": "safety",
      "configured": true
    }
  },
  "uniqueid": "00:00:00:00:00:00:00:00-00",
  "swversion": "0.000.0.00000"
}
```

5. ID指定でon/off切り替え

```sh
$ curl -s -X PUT -d "$(jo on=true)" http://<-ip-address->/api/<-user-name->/lights/1/state | jq .
[
  {
    "success": {
      "/lights/1/state/on": true
    }
  }
]

$ curl -s -X PUT -d "$(jo on=false)" http://<-ip-address->/api/<-user-name->/lights/1/state | jq .
[
  {
    "success": {
      "/lights/1/state/on": false
    }
  }
]
```

6. ID指定で色合い、輝度を変更

```sh
$ curl -s -X PUT -d "$(jo on=true bri=254 ct=366)" http://<-ip-address->/api/<-user-name->/lights/1/state | jq .
[
  {
    "success": {
      "/lights/1/state/on": true
    }
  },
  {
    "success": {
      "/lights/1/state/ct": 366
    }
  },
  {
    "success": {
      "/lights/1/state/bri": 254
    }
  }
]
```
