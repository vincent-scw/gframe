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

  get user() {
    const userStr = localStorage.getItem('user');
    return userStr && JSON.parse(userStr);
  }

  constructor() {
    this.login = this.login.bind(this);
    this.logout = this.logout.bind(this);
  }

  login(username: string) {
    let data = {name: username};
		axios.post(`${env.gameSvc}/api/user/register`, data)
			.then(res => { 
        console.log(res);
        localStorage.setItem('access_token', res.data.token);
        const decoded: any = jwt.default(res.data.token);
        const user = {
          id: decoded.sub,
          username: decoded.name
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