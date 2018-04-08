import { Component } from '@angular/core';
import {NavController, Platform} from 'ionic-angular';
import { AndroidPermissions } from '@ionic-native/android-permissions';

declare var SMS:any;

@Component({
  selector: 'page-home',
  templateUrl: 'home.html'
})
export class HomePage {

  messages:any=[ {address: "GOOD", body: "DOG"}];

  constructor(public navCtrl: NavController, public platform:Platform,
              public androidPermissions: AndroidPermissions,
              ) {

  }

  getSMS(){
    if(SMS) SMS.listSMS({}, data=>{
      setTimeout(()=>{
        console.log(data);
        this.messages=data;
      },0)
    },err=>{
      console.log(err);
    });
  }

  checkPermission()
  {
    this.androidPermissions.checkPermission
    (this.androidPermissions.PERMISSION.READ_SMS).then(
      success => {

//if permission granted
        this.ReadSMSList();
      },
      err =>{

        this.androidPermissions.requestPermission
        (this.androidPermissions.PERMISSION.READ_SMS).
        then(success=>{
            this.ReadSMSList();
          },
          err=>{
            alert("cancelled")
          });
      });
    this.androidPermissions.requestPermissions
    ([this.androidPermissions.PERMISSION.READ_SMS]);

  }
  ReadSMSList()
  {

    this.platform.ready().then((readySource) => {

      let filter = {
        box : 'inbox', // 'inbox' (default), 'sent', 'draft'
        indexFrom : 0, // start from index 0
        maxCount : 20, // count of SMS to return each time
      };

      if(SMS) SMS.listSMS(filter, (ListSms)=>{
          this.messages=ListSms
        },

        Error=>{
          alert(JSON.stringify(Error))
        });

    });
  }
}
