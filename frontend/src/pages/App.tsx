import {BrowserRouter, Navigate, Route, Routes} from "react-router-dom";
import CustomBottomNavigation from "../components/BottomNavigation/BottomNavigation.tsx";
import ErrorPage from "./ErrorPage.tsx";
import HomePage from "./Home/HomePage.tsx";
import ReferralPage from "./Referral/ReferralPage.tsx";
import StatisticPage from "./Statistic/StatisticPage.tsx";
import "../styles/root.css"
import BoostPage from "./Boost/BoostPage.tsx";
import TaskPage from "./Task/TaskPage.tsx";

const App = () => (
    <BrowserRouter>
        <div className="app">
            <Routes>
                <Route path="/" errorElement={<ErrorPage/>}>
                    <Route index element={<Navigate to="/home" replace/>}/>
                    <Route path="/home" element={<HomePage/>}/>
                    <Route path="/boost" element={<BoostPage/>}/>
                    <Route path="/task" element={<TaskPage/>}/>
                    <Route path="/referral" element={<ReferralPage/>}/>
                    <Route path="/statistic" element={<StatisticPage/>}/>
                </Route>
                <Route path="/Terrax" errorElement={<ErrorPage/>}>
                    <Route index element={<Navigate to="/home" replace/>}/>
                </Route>
            </Routes>
            <CustomBottomNavigation/>
        </div>
    </BrowserRouter>
)

export default App