import { useEffect, useState } from "react";
import { Input } from '../components/components.jsx';

const Signin = () => {
  
  const [val, setval] = useState({
    full_name: "",
    email:     "",
    password: "",
    pwd:      ""
  });

  const setter_factory = (e) => {
    setval(p => {
      return {
        ...p,
        e.target.name: e.target.value,
      }
    });
  }
  
  const SubmitData = (e) => {
    e.preventDefault();
    console.log(val);
  }

  return (
    <div className="w-full">
      <h1> Create an account? set up your information then click Goo! </h1>
      <form className="p-2 flex flex-col items-start">
        <Input name="full_name" setter={ setter_factory } Label="Full Name" type="text" placeholder="Full Name"/>
        <span className="text-red">  </span>
        <Input name="email" setter={ setter_factory } Label="Email" type="email" placeholder="Email"/>
        <Input name="password" setter={ setter_factory } Label="password" type="password" placeholder="password"/>
        <Input name="pwd" setter={ setter_factory } Label="password confirmation" type="password" placeholder="confirm password"/>
        <button onClick={ SubmitData } className="p-2 my-2 bg-blue-500 hover:bg-blue-500 transition-all text-white px-4 rounded shadow-sm">
          Go!
        </button>
      </form>
    </div>
  )
}

export default Signin;
