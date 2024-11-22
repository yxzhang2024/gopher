# CryDict Writing Tutorial

This document will guide you through the process of writing CryDict constraints for cryptographic libraries.

## 1. Function Constraints

Function constraints are applied at the function level to restrict the usage of cryptographic libraries based on the risk associated with the functions. These constraints are categorized into `DangerousFun_Constraints` and `WarningFun_Constraints`.

### Example: Custom MD5 & 3DES Functions

Suppose we have custom implementations of MD5 and 3DES in a library named `func_con/library1`:

```go
package library1

func MyMD5() {
	println("Assuming this is your own MD5 implementation")
}

func My3DES() {
	println("Assuming this is your own 3DES implementation")
}
```

The `MyMD5` function uses an outdated algorithm, which falls under `DangerousFun_Constraints`, while `My3DES` uses an algorithm that, although not entirely deprecated, should be phased out soon, falling under `WarningFun_Constraints`.

To constrain these functions, create a CryDict document corresponding to the library(in golang, package=library) named `func_con/library1`. Each document corresponds to one library and uses a blacklist approach to define function constraints:

```json
{
    "Package_Name": "func_con/library1",
    "Dangerous_Functions": [
        {
            "Name": "MyMD5",
            "Reason": "MD5 is a dangerous func"
        }
    ],
    "Warning_Functions": [
        {
            "Name": "My3DES",
            "Reason": "3DES is a warning func"
        }
    ]
}
```

Place the document at `gopher/basicRules/func_con/library1.json` and run the following command:

```
rm -rf CryDict_tutorial/func_con/scan_results
./gopher CryDict_tutorial/func_con
```

You should see that calls to `MyMD5` and `My3DES` are detected.

## 2. Parameter Constraints

Parameter constraints restrict the use of cryptographic libraries based on the parameters passed to functions.

### Example: Custom AES Function

Suppose we have a custom AES implementation in `para_con/library2`:

```go
package library2

func MyAES(key []byte) {
	println("Assuming this is your own AES implementation")
}
```

The key used with AES should not be hardcoded; thus, the first parameter (index 0) of `MyAES` should be random. The corresponding CryDict document would look like this:

```json
{
    "Package_Name": "para_con/library2",
    "Parameter_Constrains": {
        "MyAES": [
            {
                "Predicate": "RANDOM_BYTES",
                "Object": [
                    "0"
                ],
                "Reason": "AES-key should be random."
            }
        ]
    }
}
```

To write parameter constraints, specify the function name (`MyAES`), the predicate (`RANDOM_BYTES`), the object (the first parameter indexed at 0), and provide a reason explaining why the first parameter must be random.

Place the document at `gopher/basicRules/para_con/library2.json` and run:

```
rm -rf CryDict_tutorial/para_con/scan_results
./gopher CryDict_tutorial/para_con
```

You should see that the call to `MyAES` in the main function is detected.

## 3. Field Constraints

Field constraints apply restrictions to the fields of structs when using cryptographic libraries.

### Example: Custom TLS Config Struct

Suppose we have a custom TLS configuration struct in `field_con/library3`:

```go
package library3

type MyTLSConfig struct {
	InsecureSkipVerify bool
}

func (c *MyTLSConfig) Conn() {
	println("Assuming a TLS connection is established")
}
```

When establishing a TLS connection, certificate verification should not be skipped. Therefore, the field `InsecureSkipVerify` should be set to `false`. The corresponding CryDict document would be:

```json
{
    "Package_Name": "field_con/library3",
    "Struct_Constrains": {
        "MyTLSConfig": [
            {
                "Predicate": "EQ_FALSE",
                "Object": [
                    "InsecureSkipVerify"
                ],
                "Reason": "Certificate verification should not be skipped"
            }
        ]
    }
}
```

To write field constraints, specify the struct name (`MyTLSConfig`), the predicate (`EQ_FALSE`), the object (`InsecureSkipVerify`), and provide a reason explaining why the field should be set to false.

Place the document at `gopher/basicRules/field_con/library3.json` and run:

```
rm -rf CryDict_tutorial/field_con/scan_results
./gopher CryDict_tutorial/field_con
```

You should see that the misconfiguration in the main function is detected.

For more details on writing CryDict documents, refer to other files under `basicRules` and the original research paper.