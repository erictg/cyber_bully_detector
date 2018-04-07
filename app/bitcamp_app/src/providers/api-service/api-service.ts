import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import {Observable} from "rxjs/Observable";

/*
  Generated class for the ApiServiceProvider provider.

  See https://angular.io/guide/dependency-injection for more info on providers
  and Angular DI.
*/
@Injectable()
export class ApiServiceProvider {

  constructor(public http: HttpClient) {
    console.log('Hello ApiServiceProvider Provider');
  }

  baseUrl = "http://35.196.78.183:8080";

  getUser(username:string) : Observable<User>{
    return this.http.get<User>(this.baseUrl + "/rest/user/name/" + username)
  }

  getUserById(id:number) :Observable<User>{
    return this.http.get<User>(this.baseUrl + "/rest/user/id/" + id.toString())
  }

  // Name string `json:"name"`
  // IsParent bool `json:"is_parent"`
  createUser(username:string, isParent:boolean) : Observable<Object>{
    let dto = {
      "name":username,
      "is_parent":username
    };
    return this.http.post(this.baseUrl + "/rest/user", dto);
  }

  //CId int `json:"c_id"`
  // 	PId int `json:"p_id"`
  pairUsers(p_id:number, c_id:number) : Observable<Object>{
    let dto = {
      "c_id":c_id,
      "p_id":c_id
    };

    return this.http.post(this.baseUrl + "/rest/pair", dto);
  }

  // UserId int `json:"user_id"`
  // Content string `json:"content"`
  // Sent  bool `json:"sent"`
  // OtherNumber string `json:"other_number"`
  submitTextMessage(content:string, sent:boolean, phone_number:string): Observable<Object>{
    let dto = {
      "user_id": 4, //todo fix this
      "content": content,
      "sent": sent,
      "other_number": phone_number
    };

    return this.http.post(this.baseUrl + "/rest/text", dto);
  }

}

// Id 			int 	`json:"id"`
// Name 		string 	`json:"name"`
// IsParent 	bool 	`json:"is_parent"`

export class User{
  id:number;
  name:string;
  is_parent:string;
}
