import { Component } from '@angular/core';
import { NavController } from 'ionic-angular';
import {AuthServiceProvider} from "../../providers/auth-service/auth-service";
import {ApiServiceProvider, Result} from "../../providers/api-service/api-service";

@Component({
  selector: 'page-about',
  templateUrl: 'about.html'
})
export class AboutPage {

  toShow:Result;

  checkVal:string;
  constructor(public navCtrl: NavController, public auth:AuthServiceProvider, private api: ApiServiceProvider) {

  }

  check(){
    this.api.checkMean(this.checkVal).subscribe(ok =>{
      this.toShow = ok;
    },
      error1 => {
      alert("it didn't work");
      })
  }

}
