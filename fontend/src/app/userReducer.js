import { createSlice } from '@reduxjs/toolkit';
import axios from './utils/axios'
import {NotificationManager} from 'react-notifications';

export const userReducer = createSlice({
  name: 'user',
  initialState: {
    data: {}
  },
  reducers: {
    storeUser: (state, action) => {
      state.data = action.payload;
    },
  },
});

export const { storeUser } = userReducer.actions;

export const doLogin = (data, cb = () => {}) => dispatch => {
  axios.post('/user/login', data)
  .then(res => {
    dispatch(storeUser(res.data.data));
    sessionStorage.setItem("login", true)
    cb()
  })
  .catch(err => {
    NotificationManager.error(err.response ? err.response.data : 'error : something not right', '', 3000);
  })
};

export const doRegister = (data, cb = () => {}) => dispatch => {
  axios.post('/user', data)
  .then(res => {
    dispatch(storeUser(res.data.data));
    sessionStorage.setItem("login", true)
    cb()
  })
  .catch(err => {
    NotificationManager.error(err.response ? err.response.data : 'error : something not right', '', 3000);
  })
};
export const getUser = state => state;

export default userReducer.reducer;
