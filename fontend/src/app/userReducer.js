import { createSlice } from '@reduxjs/toolkit';
import axios from './utils/axios'
import {NotificationManager} from 'react-notifications';

export const userReducer = createSlice({
  name: 'user',
  initialState: {
    user: {},
  },
  reducers: {
    storeUser: (state, action) => {
      state.user = action.payload;
    },
  },
});

export const { storeUser } = userReducer.actions;

export const doLogin = data => dispatch => {
  axios.post('/user/login', data)
  .then(res => {
    dispatch(storeUser(res.data.user));
    setTimeout(() => {
      window.location = "/shift"
    }, 200);
  })
  .catch(err => {
    NotificationManager.error(err.response ? err.response.data : 'error : something not right', '', 3000);
  })
};

export const doRegister = data => dispatch => {
  axios.post('/user', data)
  .then(res => {
    dispatch(storeUser(res.data.user));
    setTimeout(() => {
      window.location = "/shift"
    }, 200);
  })
  .catch(err => {
    NotificationManager.error(err.response ? err.response.data : 'error : something not right', '', 3000);
  })
};
export const getUser = state => state.counter.value;

export default userReducer.reducer;
