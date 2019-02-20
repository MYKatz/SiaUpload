# SiaUpload

A file upload that stores your stuff on the [Sia blockchain](https://sia.tech/) - distributed and redundant. Just an experiment right now, but could potentially be useful.

## Why Sia?

Short answer: because it's cheaper. [Here's](https://siastats.info/storage_pricing) a (real-time updated) page that shows the price of storage on Sia per TB per month. The price per TB for Amazon S3 is around $20, and even low-cost competitors like [Wasabi](https://wasabi.com) can charge around $5/TB/Mo. On the other hand, the price per TB per month on Sia can be as low as 25 cents! Of course, there are reasons why it's less expensive to host on Sia: slower up/down speeds, lack of great infrastructure, etc. So this is mostly just an experiment to see how feasible it is to build a filehost completely on Sia.

## Installation

This project uses the SiaCoin API that can be accessed through localhost:9980, so you'll need to download that at [Sia.tech](https://sia.tech/get-started).

These Go Dependencies will also need to be installed:

* github.com/google/uuid
* github.com/gorilla/mux
* github.com/jinzhu/gorm
* github.com/jinzhu/gorm/dialects/sqlite

### Notes on Installation

I'm on Windows, so there were a couple things extra I needed to install. These may help if installation gets a bit screwy.

* http://tdm-gcc.tdragon.net/download - GCC compiler for windows
