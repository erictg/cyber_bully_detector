import { Component } from '@angular/core';
import { NavController } from 'ionic-angular';
import {AuthServiceProvider} from "../../providers/auth-service/auth-service";

@Component({
  selector: 'page-about',
  templateUrl: 'about.html'
})
export class AboutPage {

  constructor(public navCtrl: NavController, public auth:AuthServiceProvider) {

  }

}
