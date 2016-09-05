After [a brief chat on Facebook](https://www.facebook.com/vldid/posts/1774887999449002) regarding the recent spike of shootings in Chicago [which resulted in 91 deaths in August 2016](http://heyjackass.com/category/2016-stats/) I have promised to look at relevant data to see which claims can or cannot be backed up.

After aggregating some data from the [City of Chicago Data Portal:
Crimes - 2001 to present, 1.5GB](https://data.cityofchicago.org/Public-Safety/Crimes-2001-to-present/ijzp-q8t2) exported dataset using the `chigufi.go` Go language script the resulting data and charts are summed up in the Google spreadsheet:

### [Chicago Crime 2001 - present (Aug 25, 2016)](https://docs.google.com/spreadsheets/d/1o6gmzUn5msqEsPJyq4VxvWZ3Jj8jDHUmYZIc8vhBVj4/edit?usp=sharing)

The sheets in the document show:

0. How [the IUCR (Illinois Uniform Crime Reporting) codes](https://data.cityofchicago.org/Public-Safety/Chicago-Police-Department-Illinois-Uniform-Crime-R/c7ck-438e) were grouped for the purposed of the aggregation. IUCR codes not listed in the sheet are not used in the aggregates.
0. Graphs and the whole dataset aggregation from January 2011.
0. Graphs of the aggregation of the recent data subset from January 2013.
0. Terms of use and data disclaimer as required by the City of Chicago

Some details are off and I have no time to address them - for example, a number of homicides resulting from shootings may be higher than a reported number of murders "on the spot" in the police reports. I do not know if police retroactively repopulates these reports with newly uncovered information, and such.

I still think these charts are useful to provide context to the bare publicized number of homicides - even considering the shortcomings of this a-few-hours-hobby script.

There is intentionally no commentary or what to make out of these numbers. Interpreting the results requires a basic understanding of what is [a relationship between correlation and causation](https://xkcd.com/552/), or [what line regressions are](https://www.explainxkcd.com/wiki/index.php/1725:_Linear_Regression).

To complete the references, here is [the new Facebook post linking to the data](https://www.facebook.com/vldid/posts/1776496629288139).
