import '@styles/home-page.css'
import coinStackSvg from '@assets/coin-stack.svg'

const Counter = ({value}) => {
    return <div className="counter-container">
        <img className="counter-icon" src={coinStackSvg} alt={coinStackSvg}/>
        <div className="counter-label">{value}</div>
    </div>
}

export default Counter