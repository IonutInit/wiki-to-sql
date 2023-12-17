# Wiki to SQL

## Description

A package with a very specific purpose: extracting country data from CSV files produced from Wikipedia articles, harmonizing the data, and outputting it in CSV and SQL query formats.

The input CSVs are created using [wikitable2csv](https://github.com/gambolputty/wikitable2csv).

The reason behind all this is creating a database with various national statistics for a larger project.

## Features

- finds the country column regardless of its position in the table
- harmonizes country names to a single denominator 
- finds and parses date columns, with an option for custom date

See the _config_ file for more details.

## Improvements

Both the _country_ and the _date_ columns could be found automatically.

## Development

I did not automate this script more than it was necessary for the stated purpose, especially as I need to curate the data before passing it into the database.

However, in combination with _wikitable2csv_, end especially as it can find data such as countries and date regardless of their position, this script could be easily extended to have a just a Wikipedia link passed into it and automatically upload data into a database.