import {
  Link,
  Redirect
} from "react-router-dom"

import { useState } from 'react'
import { useDispatch } from 'react-redux'
import {doLogin} from '../app/userReducer'

export default function Login() {
  let token = localStorage.getItem("token")
  console.log("token", token)
  const dispatch = useDispatch();
  const [phone, setPhone] = useState("");
  const [password, setPassword] = useState("");

  return <>
    {/* {!token && <Redirect to="/login" />} */}
    <div className="header">
      <h4>Shift</h4>
    </div>
    <div>
      
    </div>
  </>
}