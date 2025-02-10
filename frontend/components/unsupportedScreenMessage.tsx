import IUnsupportedScreenMessageProps from "../shared/interfaces/IUnsupportedScreenMessageProps";

const UnsupportedScreenMessage: React.FC<IUnsupportedScreenMessageProps> = ({ isUnsupported }) => {
	if (isUnsupported) {
		return (
			<div className="d-flex justify-content-center align-items-center vh-100">
				<div className="text-center">
					<h4>Screen size too small</h4>
					<p>Please use a device with a larger screen to use the HabitsApp</p>
				</div>
			</div>
		);
	}

	return null;
};

export default UnsupportedScreenMessage;
