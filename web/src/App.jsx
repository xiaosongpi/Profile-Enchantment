import { BrowserRouter, Routes, Route } from "react-router-dom"
import Layout from "./components/Layout"
import Home from "./page/Home"
import User from "./page/User"
import Login from "./page/Login"

function App() {
  return(
    <BrowserRouter>
      <Routes>
        {/* Adding element element={} */}
        <Route path="/dashboard" element={<Layout />}>
          <Route path="home" element={<Home />}></Route>
          <Route path="user" element={<User />}></Route>
        </Route>
        <Route path="login" element={<Login />}></Route>
      </Routes>
    </BrowserRouter>
  )
}

export default App