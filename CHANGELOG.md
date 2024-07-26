## 0.10.0 (July 25, 2024)

NOTES:
* Moved this provider to a new namespace: `hashicorp`

## 0.8.0 (April 29, 2024)

STYLE:
* Variables passed should be passed to the last available argument in the function signature.

## 0.7.0 (April 23, 2024)

BUGS:
* Functions now prohibit the allowance of null or unknown values, with the exception of `null` and `not_null` functions.

## 0.6.0 (April 23, 2024)

FEATURES:

* **New Function:** `cidr`
* **New Function:** `cidrv4`
* **New Function:** `cidrv6`

## 0.5.0 (April 21, 2024)

FEATURES:

* change `between`, `contains`, `equal`, `not_equal` to numeric functions.

## 0.4.0 (April 5, 2024)

FEATURES:

* **New Function:** `positive`
* **New Function:** `negative`
* **New Function:** `key`
* **New Function:** `value`

## 0.3.1 (April 4, 2024)

NOTES:

* Set subcategories for all functions in the documentation

## 0.3.0 (April 4, 2024)

FEATURES:

* **New Function:** `starts_with`
* **New Function:** `ends_with`
* **New Function:** `uppercased`
* **New Function:** `lowercased`

## 0.2.0 (April 4, 2024)

FEATURES:

* **New Function:** `ip`
* **New Function:** `ipv4`
* **New Function:** `ipv6`

NOTES:

* Makefile now builds, installs, formats, lints and adds copyright headers to all files

## 0.1.2 (April 2, 2024)

* Documentation updates

## 0.1.1 (April 2, 2024)

* Documentation updates

## 0.1.0 (April 2, 2024)

NOTES:

* initial beta release
