import { BehaviorSubject } from 'rxjs';
import axios from 'axios';
import * as jwt from 'jwt-decode';
import {env} from '../services';

export class AuthService {
  userSubject = new BehaviorSubject(JSON.parse(localStorage.getItem("user") as string));

  get accessToken() {
    return localStorage.getItem('access_token');
  }

  get isLoggedin() {
    const token = this.accessToken;
    if (token == null) return false;
    const decoded: any = jwt.default(token)
    return Date.now() < decoded.exp * 1000;
  }

  constructor() {
    this.login = this.login.bind(this);
    this.logout = this.logout.bind(this);
  }

  login(username: string) {
    let data = new FormData();
		data.set("client_id", "player_api");
		data.set("client_secret", "999999");
		data.set("grant_type", "password");
		data.set("username", username);
		data.set("password", "123");
		axios.post(`${env.authSvc}/token`, data)
			.then(res => { 
        console.log(res);
        localStorage.setItem('access_token', res.data.access_token);
        const decoded: any = jwt.default(res.data.access_token);
        const user = {
          username: decoded.sub
        };
        localStorage.setItem('user', JSON.stringify(user));
        this.userSubject.next(user);
      })
			.catch(err => console.error(err));
  }

  logout() {
    localStorage.removeItem('user');
    localStorage.removeItem('access_token');
    this.userSubject.next(null);
  }
}