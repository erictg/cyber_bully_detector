import { Component } from '@angular/core';
import { Platform } from 'ionic-angular';
import { StatusBar } from '@ionic-native/status-bar';
import { SplashScreen } from '@ionic-native/splash-screen';

import { TabsPage } from '../pages/tabs/tabs';
import {FCM} from "@ionic-native/fcm";
import {ApiServiceProvider} from "../providers/api-service/api-service";

@Component({
  templateUrl: 'app.html'
})
export class MyApp {
  rootPage:any = TabsPage;

  constructor(platform: Platform, statusBar: StatusBar, splashScreen: SplashScreen, private fcm: FCM, private api: ApiServiceProvider) {
    platform.ready().then(() => {
      // Okay, so the platform is ready and our plugins are available.
      // Here you can do any higher level native things you might need.
      statusBar.styleDefault();
      splashScreen.hide();
      this.fcm.subscribeToTopic('all');
      this.fcm.getToken().then(token => {
        api.updateFCM(token);
      });
      this.fcm.onNotification().subscribe(data => {
        alert('message received');
        if(data.wasTapped) {
          console.info("Received in background");
        } else {
          console.info("Received in foreground");
        }
      });
      this.fcm.onTokenRefresh().subscribe(token => {
        api.updateFCM(token);
      });
    });
  }
}
