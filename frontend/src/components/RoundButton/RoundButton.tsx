import '../../styles/round-button.css'

const RoundButton = ({onClick, imgSrc}) => (
    <button onClick={onClick} className="round-button">
        <img className="round-button-img" src={imgSrc} alt={imgSrc}/>
    </button>
)

export default RoundButton