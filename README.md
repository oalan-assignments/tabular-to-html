## Why?
 
We are interested in your skills as a developer. As part of our assessment, we want to see your code.
 
## Instructions
 
In this archive, you'll find two files, Workbook2.csv and Workbook2.prn, which need to be displayed by the software you deliver in HTML format.
 
This repository is created specially for you, so you can push anything you like. Please update this README to provide instructions, notes and/or comments for us.
 
 
The solution has to be implemented in Go and we expect this to be done within a week.

## My Comments

## To run
go run workbook_to_html.go

and then please try
[Prn file](http://localhost:8080/Workbook2.prn) 
or 
[Csv file](http://localhost:8080/Workbook2.csv) 

## What I would do next?
More elaborate parsing:
* Line/header skipping
* Handling arbitrary schema
* Handling arbitrary encoding

As you may notice from return statement of prn parser (error as nil), there is not much error
handling in prn parsing. I would be a bit more defensive with the help of unit
tests as I did with csv parsing.

I would spend more time to learn idiomatic go programming. 
I feel like I had to go through things superficially.

I thank you in advance! :)
