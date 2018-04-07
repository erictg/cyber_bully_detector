import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import {User} from "../api-service/api-service";
import {Storage} from "@ionic/storage";
import {Observable} from "rxjs/Observable";

/*
  Generated class for the StorageServiceProvider provider.

  See https://angular.io/guide/dependency-injection for more info on providers
  and Angular DI.
*/
@Injectable()
export class StorageServiceProvider {

  constructor(public storage:Storage) {
  }

  saveUser(user:User){
    this.storage.set("user", JSON.stringify(user)).then();
  }

  clearUser(){
    this.storage.set("user", "").then();
  }

  getUser() : Promise<User>{
    return this.storage.get("user");

  }

}
