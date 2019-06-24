// in src/authProvider.js
import { AUTH_LOGIN } from 'react-admin';

export default (type, params) => {
    if (type === AUTH_LOGOUT) {
        this.logout()
    }
    return Promise.resolve();
}