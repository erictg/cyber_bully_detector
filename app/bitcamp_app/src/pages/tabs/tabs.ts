import { Component } from '@angular/core';

import { AboutPage } from '../about/about';
import { ContactPage } from '../contact/contact';
import { HomePage } from '../home/home';
import {AuthServiceProvider} from "../../providers/auth-service/auth-service";
import {ApiServiceProvider, User} from "../../providers/api-service/api-service";
import {HttpErrorResponse} from "@angular/common/http";

@Component({
  templateUrl: 'tabs.html'
})
export class TabsPage {
  username:string = '';
  tab1Root = HomePage;
  tab2Root = AboutPage;
  tab3Root = ContactPage;


  parent:boolean = false;
  constructor(private auth: AuthServiceProvider, private api:ApiServiceProvider) {

  }


  createUser(){
    this.api.createUser(this.username, this.parent).subscribe(value => {
      if(typeof (value) == typeof (HttpErrorResponse)){
        console.log('in typeif');
        if((value as HttpErrorResponse).status != 200){
          console.log("in check");
        }
      }else{
        this.auth.login(this.username);
      }
    })
  }
}
