import React, {useEffect, useState} from "react";
import RoundButton from "../../components/RoundButton/RoundButton.tsx";
import coinImg from '../../assets/coin.svg'
import '@styles/home-page.css'
import Counter from "@/pages/Home/components/Counter.tsx";

const HomePage = () => {
    useEffect(() => {

    }, []);
    const [counter, setCounter] = useState(0);

    const handleClick = (e: React.MouseEvent<HTMLButtonElement, MouseEvent>) => {
        setCounter(prev => prev + 1);

        const button = e.currentTarget;
        const rect = button.getBoundingClientRect();
        const offsetX = e.clientX - rect.left;
        const offsetY = e.clientY - rect.top;

        const fadeOutText = document.createElement('div');
        fadeOutText.classList.add('fade-out');
        fadeOutText.textContent = '+1';
        fadeOutText.style.left = `${offsetX}px`;
        fadeOutText.style.top = `${offsetY}px`;
        button.appendChild(fadeOutText);

        setTimeout(() => {
            button.removeChild(fadeOutText);
        }, 1000);
    };

    return <div className="container">
        <div className="counter-section">
            <Counter value={counter}/>
        </div>
        <div className="button-section">
            <RoundButton imgSrc={coinImg} onClick={handleClick}/>
        </div>
        <div className="empty-section"></div>
    </div>
}

export default HomePage;