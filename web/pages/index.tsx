import 'bootstrap/dist/css/bootstrap.css'
import { useFormik } from 'formik';
import {useEffect, useState} from "react";


export async function getServerSideProps({req}) {
    let isAuthorization : boolean = false
    if(req.cookies["token"]){
        const res = await fetch("http://127.0.0.1:3001/api/authorization",{
            withCredentials: true,
            headers: {
                Cookie: req.headers.cookie
            }
        })
        const json = await  res.json()
        isAuthorization = json.status
    }
    return {
        props: {
            isAuthorization:isAuthorization
        },
    }
}

export default function Index({isAuthorization}) {
    const [username, setUsername] = useState()
    const [password, setPassword] = useState()


    const [isLogged, setIsLogged] = useState(isAuthorization);
    const [errorMsg, setErrorMsg] = useState(null);

    const logout = async () => {
        try{
            const d = await fetch("/api/logout")
            setIsLogged(false)
        }catch (err){
            console.log(err)
        }
    }

    const authorization = async (event) => {
        event.preventDefault()
        fetch("/api/authentication", {
            method: "POST",
            body: JSON.stringify({username:username , password:password}),
            headers:{"Content-Type": "application/json"}
        })
            .then(data => data.json())
            .then(json => {
                if(json.status === true){
                    setIsLogged(true)
                    setErrorMsg(null)
                }else {
                    setErrorMsg(json.msg)
                }
            })
            .catch(err => {
                console.log(err)
            })
    }


    // const formik = useFormik({
    //         initialValues: {
    //            username: "", password: "",
    //         }, onSubmit: (values) => {
    //             console.log(values.username +"  /  "+  values.password)
    //             setAccount({username: values.username , password: values.password})
    //             // alert(JSON.stringify(values, null, 2));
    //             authorization()
    //         },
    // });

    // useEffect(() => {
        if(isLogged === true){
            return(
                <div>
                    <div className="Auth-form-container">
                        <div className="form-group mt-3 m-2">
                            <span>Authenticated already</span>
                            <button type="submit" className="m-2 btn btn-danger" onClick={logout}>
                                Logout
                            </button>
                        </div>
                    </div>
                </div>
            )
        }
    // })


    return (
        <div className="Auth-form-container">
            <form className="Auth-form" onSubmit={authorization} method="POST">
                <div className="Auth-form-content">
                    <h3 className="Auth-form-title">Sign In</h3>
                    <div className="form-group mt-3">
                        {
                            errorMsg &&
                            <div className="alert alert-danger" role="alert">
                                {errorMsg}
                            </div>
                        }
                    </div>
                    <div className="form-group mt-3">
                        <label>USERNAME</label>
                        <input
                            id="username"
                            name="username"
                            type="text"
                            className="form-control mt-1"
                            placeholder="Enter username"
                            onChange={e => setUsername(e.target.value)}
                        />
                    </div>
                    <div className="form-group mt-3">
                        <label>PASSWORD</label>
                        <input
                            id="password"
                            name="password"
                            type="password"
                            className="form-control mt-1"
                            placeholder="Enter password"
                            onChange={e => setPassword(e.target.value)}
                        />
                    </div>
                    <div className="d-grid gap-2 mt-3">
                        <div className="row">
                            <div className="col-sm">
                                <label>OAUTH2</label>
                            </div>
                        </div>
                        <div className="row">
                            <div className="col-sm">
                                <button  className="btn btn-light">
                                    Google
                                </button>
                            </div>
                            <div className="col-sm">
                                <button className="btn btn-primary text-white">
                                    Facebook
                                </button>
                            </div>
                            <div className="col-sm">
                                <button className="btn btn-dark">
                                    Github
                                </button>
                            </div>
                        </div>
                    </div>
                    <div className="d-grid gap-2 mt-3">
                        <button type="submit" className="btn btn-primary">
                            Submit
                        </button>
                    </div>
                </div>
            </form>
        </div>
    )
}
