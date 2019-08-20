import axios from 'axios';
import authService from './auth.service';
import env from './environment';

class PlayerReceptionService {
  in() {
    return axios.post(`${env.gameSvc}/api/user/in`, null, {
      headers: {'Authorization': `Bearer ${authService.accessToken}`}
    })
  }

  out() {
    return axios.post(`${env.gameSvc}/api/user/out`, null, {
      headers: {'Authorization': `Bearer ${authService.accessToken}`}
    })
  }
}

const playerReceptionService = new PlayerReceptionService();
export default playerReceptionService;