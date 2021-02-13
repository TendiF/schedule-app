import {
  Link,
  Redirect
} from "react-router-dom"

import { useSelector } from 'react-redux';
import { useState } from 'react'
import { useDispatch } from 'react-redux'
import {getUser} from '../app/userReducer'

export default function Login() {
  const {user} = useSelector(getUser);
  
  return <>
    {!user.data.id && <Redirect to="/login" />}
    <div className="header">
      <h4>Shift {user.data.name}</h4>
    </div>
    <div>
      
    </div>
  </>
}