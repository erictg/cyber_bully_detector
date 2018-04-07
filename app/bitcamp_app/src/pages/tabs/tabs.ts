import { Component } from '@angular/core';

import { AboutPage } from '../about/about';
import { ContactPage } from '../contact/contact';
import { HomePage } from '../home/home';
import {AuthServiceProvider} from "../../providers/auth-service/auth-service";

@Component({
  templateUrl: 'tabs.html'
})
export class TabsPage {
  username:string = '';
  tab1Root = HomePage;
  tab2Root = AboutPage;
  tab3Root = ContactPage;

  constructor(private auth: AuthServiceProvider) {

  }
}
