import { AuthService } from "./auth.service";
import { GameService } from "./game.service";
import { environment } from "./environment";

const authService = new AuthService();
export {authService};

const gameService = new GameService();
export {gameService};

const env = new environment();
export {env};