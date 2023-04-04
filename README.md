# restaurant-recomm

Restaurant recommender

Problem: https://docs.google.com/document/d/1Q2sP8O1NXSa1xcNZshscydi1I-KbD_UekTxxrbcFONs/edit

###  Design
Each folder of src contain one package

#### Loader package
This package loads and parses [data.csv](https://github.com/tiwarishub/restaurant-recomm/blob/main/data.csv). 
 
#### Models
This package basically holds all models arounds our resturant recommender. This include restaurant, cuisines and user.
 
#### Rules
This package contains all rules which are used by restaurant recommender. In this, we have rules.go which basically contains one interface which are implemented by different types of rules we have.


#### Recommender
This is main engine for our restaurant recommender. This package basically create rules using rule package and apply those rules based on user and restaurant data.



### Testing

To test this code just run below code from root folder of the repo
```
go run src/main/main.go
```

Based on current user set in [main.go](https://github.com/tiwarishub/restaurant-recomm/blob/main/src/main/main.go) and data for restaurants present in [data.csv](https://github.com/tiwarishub/restaurant-recomm/blob/main/data.csv). It output is
```
ID2
ID5
ID17
ID3
ID6
ID7
ID4
ID1
ID19
ID9
ID10
ID8
ID11
ID12
ID13
ID14
ID15
ID16
ID18
ID20
ID21
ID22
ID23
```

To understand this output, let analyse the [data.explain](https://github.com/tiwarishub/restaurant-recomm/blob/main/data.explain) which basically explain the entries in data.csv file. Here is below explained output based on  rules defined in the problem
```
ID2  -> Rule 1 "Featured restaurants of primary cuisine and primary cost bracket."
ID5  -> Rule 2 "All restaurants of Primary cuisine, primary cost bracket with rating >= 4"
ID17 -> Rule 3 "All restaurants of Primary cuisine, secondary cost bracket with rating >= 4.5"
ID3  -> Rule 4 "All restaurants of secondary cuisine, primary cost bracket with rating >= 4.5"
ID6  -> Rule 4 "All restaurants of secondary cuisine, primary cost bracket with rating >= 4.5"
ID7  -> Rule 5 "Top 4 newly created restaurants by rating"
ID4  -> Rule 5  
ID1  -> Rule 5
ID19 -> Rule 5
ID9  -> Rule 6 "All restaurants of Primary cuisine, primary cost bracket with rating < 4"
ID10 -> Rule 8 "All restaurants of secondary cuisine, primary cost bracket with rating < 4.5"
ID8  -> Rule 9 "All restaurants of any cuisine, any cost bracket"
ID11 -> Rule 9
ID12 -> Rule 9
ID13 -> Rule 9
ID14 -> Rule 9
ID15 -> Rule 9
ID16 -> Rule 9
ID18 -> Rule 9
ID20 -> Rule 9
ID21 -> Rule 9
ID22 -> Rule 9
ID23 -> Rule 9
```
**Note**: No combination of sample user defined main.go and data in data.csv satified rule 7. So no explicit entry for rule7


 
