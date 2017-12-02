* The below are the list of the URI available in this service
* Please check below urls and other information in order to execute url

1. Add a trainer
url: http://localhost:9392/AddTrainer

Method Type: POST

Headers: Content-Type=application/json
Body: 
{
	"name":"Muruga",
	"dob":"14-Aug",
	"gender":"male",
	"address":"Munipalle , Ponnur Mandla, Guntur District, AP, India, 522316",
	"email":"jitenp@outlook.com",
	"web":"www.linkedin.com/jpalaparthi",
	"contact":"9618558500",
	"skills":["c","c++","java","golang"],
	"status":"active"
	
	ID        bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Name      string        `json:"name" bson:"name"`
	DOB       string        `json:"dob" bson:"dob"`
	Gender    string        `json:"gender" bson:"gender"`
	Address   string        `json:"address" bson:"address"`
	Email     string        `json:"email" bson:"email"`
	Web       string        `json:"web" bson:"web"`
	Contact   string        `json:"contact" bson:"contact"`
	Skills    []string      `json:"roles" bson:"skills"`
	Status    string        `json:"status" bson:"status"`
	TimeStamp string        `json:"timestamp" bson:"timestamp"`
	
	
	
}

OnSuccess:

{
  "message": "Movie successfully added",
  "success": true,
  "status": 200,
  "trace": "58c77023-52fd-fc07-2182-654f-163f5f0f"
}

OnFailure:
{
  "message": "failure reason",
  "success": false,
  "status": 400,
  "trace": "failure reason"
}

-------------------------------------------------------------
2. FetchAllMovies
url: http://localhost:9392//FetchAllMovies"

Method Type: GET

Headers: 

OnSuccess:

[
  {
    "_id": "58c77023307bf5e715c7df9f",
    "batchno": "1",
    "movieid": "58c77023-52fd-fc07-2182-654f-163f5f0f",
    "title": "Singam1",
    "wiki": ""
  }
]

OnFailure:
{
  "message": "no data found",
  "success": false,
  "status": 400,
  "trace": "no data found"
}

-------------------------------------------------------------
3. GetMovieList : Batch number is mandatory
url: http://localhost:9392/GetMovieList?batchno=1

Method Type: GET

Headers: 

OnSuccess:

[
  {
    "_id": "58c77023307bf5e715c7df9f",
    "batchno": "1",
    "movieid": "58c77023-52fd-fc07-2182-654f-163f5f0f",
    "title": "Singam1",
    "wiki": ""
  }
]

OnFailure:

{
  "message": "no data available",
  "success": false,
  "status": 400,
  "trace": "no data available"
}

-------------------------------------------------------------
4. GetMovieByMovieID : Movie id is mandatory
url: http://localhost:9392/GetMovieByMovieID?movieid=1

Method Type: GET

Headers: 

OnSuccess:

{
  "_id": "58c77023307bf5e715c7df9f",
  "batchno": "1",
  "movieid": "58c77023-52fd-fc07-2182-654f-163f5f0f",
  "title": "Singam1",
  "wiki": ""
}

OnFailure:

{
  "message": "no data available",
  "success": false,
  "status": 400,
  "trace": "no data available"
}

-------------------------------------------------------------
5. GetMoviesByKeyword : Key is db json key and valye is db json value.key and value are mandatory
url: http://localhost:9392/GetMoviesByKeyword?key=batchno&value=1
url: http://localhost:9392/GetMoviesByKeyword?key=title&value=Singam1

Method Type: GET

Headers: 

OnSuccess:

[
  {
    "_id": "58c77023307bf5e715c7df9f",
    "batchno": "1",
    "movieid": "58c77023-52fd-fc07-2182-654f-163f5f0f",
    "title": "Singam1",
    "wiki": ""
  }
]

OnFailure:

{
  "message": "no data available",
  "success": false,
  "status": 400,
  "trace": "no data available"
}

-------------------------------------------------------------
6. UploadImageForMovie : movieid,title and uploadfile are mandatory.All are form-data fields

url: http://localhost:9392//UploadImageForMovie

Method Type: POST

Headers: form-data
movieid:"58c77023-52fd-fc07-2182-654f-163f5f0f"
title:"Singam1"
relevance:45
option:true
uploadfile:(file with the path)

OnSuccess:

Description: File is nude and hence sent to quarantine folder.

{
  "message": "flagged",
  "success": true,
  "status": 400,
  "trace": "58c7732e-9a62-1d72-9566-c74d-10037c4d"
}

Description: File is short and hence sent to invalid folder.
{
  "message": "short",
  "success": true,
  "status": 400,
  "trace": "58c77427-7bbb-0407-d1e2-c649-81855ad8"
}

Description: File is verified as okay. and hence send to full folder and asynchly medium folder
{
  "message": "Pic successfully uploaded",
  "success": true,
  "status": 200,
  "trace": "1"
}

OnFailure:

{
  "message": "failure message",
  "success": false,
  "status": 400,
  "trace": "failure message"
}

-------------------------------------------------------------
7. GetMoviesByKeyword : fetches 4 images(configarable).min ,max and movieid are mandatory querystrings
url: http://localhost:9392/GetImagesByDelta?min=3&max=4&movieid=58c77023-52fd-fc07-2182-654f-163f5f0f

Method Type: GET

Headers: 

OnSuccess:

Description:
 
* Checks number of records to fetch and checks current highest file number.
* In the below result, 23 Files are there , min is 3 and max is 4 and 4 to fetch, hence
* It fetches 23,22,21 and 20 file records respectively


[
  {
    "_id": "58c77517307bf5e715c7dfb8",
    "filetype": ".jpg",
    "movieid": "58c77023-52fd-fc07-2182-654f-163f5f0f",
    "option": "true",
    "picid": "23",
    "picpath": "uploads/58c77023-52fd-fc07-2182-654f-163f5f0f/full/23",
    "picstatus": "valid",
    "relevence": "45.6",
    "status": "Active",
    "timestamp": "2017-03-14 10:14:07.620121013 +0530 IST",
    "title": "Singam1"
  },
  {
    "_id": "58c77517307bf5e715c7dfb7",
    "filetype": ".jpg",
    "movieid": "58c77023-52fd-fc07-2182-654f-163f5f0f",
    "option": "true",
    "picid": "22",
    "picpath": "uploads/58c77023-52fd-fc07-2182-654f-163f5f0f/full/22",
    "picstatus": "valid",
    "relevence": "45.6",
    "status": "Active",
    "timestamp": "2017-03-14 10:14:07.196677471 +0530 IST",
    "title": "Singam1"
  },
  {
    "_id": "58c77516307bf5e715c7dfb6",
    "filetype": ".jpg",
    "movieid": "58c77023-52fd-fc07-2182-654f-163f5f0f",
    "option": "true",
    "picid": "21",
    "picpath": "uploads/58c77023-52fd-fc07-2182-654f-163f5f0f/full/21",
    "picstatus": "valid",
    "relevence": "45.6",
    "status": "Active",
    "timestamp": "2017-03-14 10:14:06.711848627 +0530 IST",
    "title": "Singam1"
  },
  {
    "_id": "58c77516307bf5e715c7dfb5",
    "filetype": ".jpg",
    "movieid": "58c77023-52fd-fc07-2182-654f-163f5f0f",
    "option": "true",
    "picid": "20",
    "picpath": "uploads/58c77023-52fd-fc07-2182-654f-163f5f0f/full/20",
    "picstatus": "valid",
    "relevence": "45.6",
    "status": "Active",
    "timestamp": "2017-03-14 10:14:06.283652665 +0530 IST",
    "title": "Singam1"
  }
]
-----

[
  {
    "_id": "58c77517307bf5e715c7dfb8",
    "filetype": ".jpg",
    "movieid": "58c77023-52fd-fc07-2182-654f-163f5f0f",
    "option": "true",
    "picid": "23",
    "picpath": "uploads/58c77023-52fd-fc07-2182-654f-163f5f0f/full/23",
    "picstatus": "valid",
    "relevence": "45.6",
    "status": "Active",
    "timestamp": "2017-03-14 10:14:07.620121013 +0530 IST",
    "title": "Singam1"
  },
  {
    "_id": "58c77517307bf5e715c7dfb7",
    "filetype": ".jpg",
    "movieid": "58c77023-52fd-fc07-2182-654f-163f5f0f",
    "option": "true",
    "picid": "22",
    "picpath": "uploads/58c77023-52fd-fc07-2182-654f-163f5f0f/full/22",
    "picstatus": "valid",
    "relevence": "45.6",
    "status": "Active",
    "timestamp": "2017-03-14 10:14:07.196677471 +0530 IST",
    "title": "Singam1"
  },
  {
    "_id": "58c7750e307bf5e715c7dfa4",
    "filetype": ".jpg",
    "movieid": "58c77023-52fd-fc07-2182-654f-163f5f0f",
    "option": "true",
    "picid": "3",
    "picpath": "uploads/58c77023-52fd-fc07-2182-654f-163f5f0f/full/3",
    "picstatus": "valid",
    "relevence": "45.6",
    "status": "Active",
    "timestamp": "2017-03-14 10:13:58.334115821 +0530 IST",
    "title": "Singam1"
  },
  {
    "_id": "58c7750d307bf5e715c7dfa3",
    "filetype": ".jpg",
    "movieid": "58c77023-52fd-fc07-2182-654f-163f5f0f",
    "option": "true",
    "picid": "2",
    "picpath": "uploads/58c77023-52fd-fc07-2182-654f-163f5f0f/full/2",
    "picstatus": "valid",
    "relevence": "45.6",
    "status": "Active",
    "timestamp": "2017-03-14 10:13:57.189272921 +0530 IST",
    "title": "Singam1"
  }
]
Description:

* Checks number of records to fetch and checks current highest file number.
* In the below result, 23 Files are there , min is 4 and max is 21 and 4 to fetch, hence
* It fetches 23,22,3 and 2 file records respectively

OnFailure:

{
  "message": "no data available",
  "success": false,
  "status": 400,
  "trace": "no data available"
}

-------------------------------------------------------------
8. GetMoviesByKeyword : movidid,picid and imagesize(full or medium) querystrings are mandatory

url: http://localhost:9392/GetFullImageByID?movieid=58c77023-52fd-fc07-2182-654f-163f5f0f&picid=22&imagesize=full

Method Type: GET

Headers: 

OnSuccess:

* Returns an image based. 
* If imagesize=full returns full image.
* If imagesize=medium returns medium image.

OnFailure:

{
  "message": "no data available",
  "success": false,
  "status": 400,
  "trace": "no data available"
}

-------------------------------------------------------------