import axios from 'axios';
import authService from './auth.service';

class PlayerReceptionService {
  in() {
    axios.post('http://localhost:8080/user/in', null, {
      headers: {'Authorization': `Bearer ${authService.accessToken}`}
    })
  }

  out() {
    axios.post('http://localhost:8080/user/out', null, {
      headers: {'Authorization': `Bearer ${authService.accessToken}`}
    })
  }
}

const playerReceptionService = new PlayerReceptionService();
export default playerReceptionService;