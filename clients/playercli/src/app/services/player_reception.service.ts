import axios from 'axios';
import authService from './auth.service';

class PlayerReceptionService {
  in() {
    return axios.post('http://localhost:8441/api/user/in', null, {
      headers: {'Authorization': `Bearer ${authService.accessToken}`}
    })
  }

  out() {
    return axios.post('http://localhost:8441/api/user/out', null, {
      headers: {'Authorization': `Bearer ${authService.accessToken}`}
    })
  }
}

const playerReceptionService = new PlayerReceptionService();
export default playerReceptionService;