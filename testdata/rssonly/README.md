# rss feed processing rules
Most rss feeds are broken and inconsistent, so we need some rules to deal with these rss feeds in a consistent way

# fields
Here is the standard field level processing rules

## Published date

1. When the published date is not found, then return nil. This will make the consumer to use Processed Date
2. When the published date is failed parsing, then return nil. This will make the consumer to use Processed Date
3. When the published date is missing time, then return the date with time set at midnight. This will make the consumer to use Processed Date

> We only load data of a particular source link into the consumer only once so if the published date is missing, we can use processed date as published date



