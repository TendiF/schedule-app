import { configureStore } from '@reduxjs/toolkit';
import counterReducer from '../features/counter/counterSlice';
import userReducer from './userReducer';
import shiftReducer from './shiftReducer';
export default configureStore({
  reducer: {
    counter: counterReducer,
    user: userReducer,
    shift: shiftReducer,
  },
});
