import { createSlice } from '@reduxjs/toolkit';
import axios from './utils/axios'
import {NotificationManager} from 'react-notifications';

export const shiftReducer = createSlice({
  name: 'user',
  initialState: {
    data: []
  },
  reducers: {
    storeShift: (state, action) => {
      state.data = action.payload;
    },
  },
});

export const { storeShift } = shiftReducer.actions;

export const getAxiosShift = (data, cb = () => {}) => dispatch => {
  axios.get('/shift', {
    params: data
  }, { headers: { "Content-Type": "application/json" } })
  .then(res => {
    dispatch(storeShift(res.data.data));
    cb()
  })
  .catch(err => {
    NotificationManager.error(err.response ? err.response.data : 'error : something not right', '', 3000);
  })
};

export const getShift = state => state;

export default shiftReducer.reducer;
