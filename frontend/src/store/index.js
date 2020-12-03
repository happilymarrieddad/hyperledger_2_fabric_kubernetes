import { createStore } from 'vuex';

export default createStore({
  state: {
      token: false,
  },
  mutations: {
      setToken(state, token) {
        state.token = token;
      }
  },
  actions: {
      async login({ commit }, { username = '', password = '' }) {
          return new Promise((resolve) => {
            // TODO: make a call to the server to authorize the request

            commit('setToken', '1234');

            // TODO: perhaps return a token
            return resolve(['', null]);
          })
      },
      async hasPermissionForRoute({}, { route = '' }) {
        // TODO: Check
        return new Promise((resolve) => {
            return resolve([true, null]);
        })
      }
  },
  getters: {
    isAuthenticated: state => {
        // TODO: check the token
        return true;
        //return state.auth;
    }
  },
  modules: {
  },
});
