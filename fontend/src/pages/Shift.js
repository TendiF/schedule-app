import {
  Link,
  Switch,
  Route,
  Redirect,
  useParams,
  useRouteMatch
} from "react-router-dom"

import { useSelector } from 'react-redux';
import { useState, useEffect } from 'react'
import { useDispatch } from 'react-redux'
import { getUser } from '../app/userReducer'
import { getAxiosShift, getShift } from '../app/shiftReducer'

export default function Shift() {
  const {user} = useSelector(getUser);
  let { path, url } = useRouteMatch();
  const [userData, setUser] = useState({});
  
  let login = sessionStorage.getItem("login")

  useEffect(() => {
    if(user.data){
      setUser(user.data)
    }
  }, [user, userData])

  return <>
    {!login && <Redirect to="/login" />} 
    <div className="header">
      <h4>Shift {userData.name}</h4>
    </div>
    <Switch>
        <Route exact path={path}>
         <ul>
            <li>
              <Link to={`${url}/form`}>Form</Link>
            </li>
            <li>
              <Link to={`${url}/1`}>detail </Link>
            </li>
          </ul>
          <ListShift user={userData}/>
        </Route>
        <Route path={`${path}/form`}>
          <FormShift />
        </Route>
        <Route path={`${path}/:id`}>
          <DetailShift />
        </Route>
    </Switch>
  </>
}

function DetailShift(){
  let { id } = useParams();
  return <div>
    Detail Shift
  </div>
}

function FormShift(){
  let { id } = useParams();
  return <div>
    Form Shift
  </div>
}

function ListShift(props){
  const dispatch = useDispatch();
  const {shift} = useSelector(getShift);

  useEffect(() => {
    dispatch(getAxiosShift({id_user : props.user.id}))
  }, [dispatch])

  return <div>
    {Array.isArray(shift.data) && !shift.data.length && "empty data"}
    {Array.isArray(shift.data) && shift.data.map((v,i) => {
      return <div key={v.start_date} className="shiftCard">
        <h4>Name</h4>
        <StartEnd start_date={v.start_date} end_date={v.end_date}/>
      </div>
    })}
  </div>
}

let StartEnd = (props) => {
  let {start_date, end_date} = props
  let comp = null
  if(new Intl.DateTimeFormat('en-GB', { dateStyle: 'short'}).format(new Date(end_date)) === new Intl.DateTimeFormat('en-GB', { dateStyle: 'short'}).format(new Date(start_date))){
    comp = <p>
      {new Intl.DateTimeFormat('en-GB', { dateStyle: 'short', timeStyle: 'short' }).format(new Date(start_date))} - {new Intl.DateTimeFormat('en-GB', { timeStyle: 'short' }).format(new Date(end_date))}
      </p>
  }else {
    comp = <p>
    {new Intl.DateTimeFormat('en-GB', { dateStyle: 'short', timeStyle: 'short' }).format(new Date(start_date))} - {new Intl.DateTimeFormat('en-GB', { dateStyle: 'short', timeStyle: 'short' }).format(new Date(end_date))}
    </p>
  }

  return <>{comp}</>
}