import { Injectable } from '@angular/core';
import {StorageServiceProvider} from "../storage-service/storage-service";
import {ApiServiceProvider, User} from "../api-service/api-service";

/*
  Generated class for the AuthServiceProvider provider.

  See https://angular.io/guide/dependency-injection for more info on providers
  and Angular DI.
*/
@Injectable()
export class AuthServiceProvider {


  public isLoggedIn = false;
  public user:User = null;
  constructor(private storageService:StorageServiceProvider, private api:ApiServiceProvider) {
    console.log('Hello AuthServiceProvider Provider');
    storageService.getUser().then(value => {
      console.log(value);
      if(value != null){
        this.user = value;
        this.isLoggedIn = true;
      }
    });
  }

  // Login a user
  // Normally make a server request and store
  // e.g. the auth token
  login(username:string) : void {
    this.api.getUser(username).subscribe(user=>{
      if(user == null){
        this.isLoggedIn = false;
      }else{
        this.user = user;
      }
    });
    this.isLoggedIn = true;
  }

  // Logout a user, destroy token and remove
  // every information related to a user
  logout() : void {
    this.storageService.clearUser();
    this.isLoggedIn = false;
  }


}
