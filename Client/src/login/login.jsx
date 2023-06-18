import { useEffect, useState } from "react";
import { Input } from '../components/components.jsx';

const Login = () => {
  
  const [val, setval] = useState({
    email:     "",
    password:  "",
  });
  
  const setter_factory = (e) => {
    setval(p => {
      return {
        ...p,
        e.target.name: e.target.value,
      }
    });
  }

  return (
    <div className="w-full">
      <h1> Login to your account! </h1>
      <form className="p-2 flex flex-col items-start">
        <Input name="email"    setter={ setter_factory } Label="Email" type="email" placeholder="Email"/>
        <Input name="password" setter={ setter_factory } Label="password" type="password" placeholder="password"/>
        
        <button className="p-2 my-2 bg-blue-500 hover:bg-blue-500 transition-all text-white px-4 rounded shadow-sm">
          Login!
        </button>

      </form>
    </div>
  )
}

export default Signin;
