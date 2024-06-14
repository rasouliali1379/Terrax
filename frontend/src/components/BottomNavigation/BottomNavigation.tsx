import * as React from 'react';
import Box from '@mui/material/Box';
import BottomNavigation from '@mui/material/BottomNavigation';
import BottomNavigationAction from '@mui/material/BottomNavigationAction';
import BarChartIcon from '@mui/icons-material/BarChart';
import Diversity1Icon from '@mui/icons-material/Diversity1';
import {useNavigate} from "react-router-dom";
import {Bolt, Checklist, Home} from "@mui/icons-material";


export default function CustomBottomNavigation() {
    const [value, setValue] = React.useState('/home');
    const navigate = useNavigate();

    return (
        <Box sx={{width: '100%', position: 'fixed', bottom: 0, left: 0, right: 0, display: 'flex'}}>
            <BottomNavigation
                sx={{width: '100%', borderRadius: '10px', margin: '8px', boxShadow: "2px 2px 6px #00000040"}}
                showLabels
                value={value}
                onChange={(_, newValue) => {
                    setValue(newValue);
                    navigate(newValue)
                }}
            >
                <BottomNavigationAction sx={{minWidth: "50px"}} label="Referral" icon={<Diversity1Icon/>}
                                        value="/referral"/>
                <BottomNavigationAction sx={{minWidth: "50px"}} label="Task" icon={<Checklist/>} value="/task"/>
                <BottomNavigationAction sx={{minWidth: "50px"}} label="Home" icon={<Home/>} value="/home"/>
                <BottomNavigationAction sx={{minWidth: "50px"}} label="Boost" icon={<Bolt/>} value="/boost"/>
                <BottomNavigationAction sx={{minWidth: "50px"}} label="Statistic" icon={<BarChartIcon/>}
                                        value="/statistic"/>
            </BottomNavigation>
        </Box>
    );
}