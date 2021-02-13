import {
  Link,
} from "react-router-dom"

import { useState } from 'react'
import { useDispatch } from 'react-redux'
import {doLogin} from '../app/userReducer'

export default function Login() {
  const dispatch = useDispatch();
  const [name, setName] = useState("");

  return <>
    <div className="header">
      <h4>Login</h4>
    </div>
    <div>
      <div className="login">
        <div>
          <p>Name</p>
          <div className="input" style={{alignItems:'center'}} >
            <input value={name} onChange={e => setName(e.target.value)} placeholder="Name" style={{height:'30px', marginLeft:'10px'}}/>
          </div>
        </div>
      <button onClick={() => dispatch(doLogin({name}))} style={{marginTop: '15px'}} className="button-primary">Login</button>
      <Link style={{marginTop:'10px'}} to="register">Don't have account ? <span style={{color:'blue'}}>go to register</span> </Link>
      </div>
    </div>
  </>
}