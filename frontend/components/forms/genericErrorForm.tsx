import React from "react";
import IGenericErrorForm from "../../shared/interfaces/IGenericErrorForm";

const GenericErrorForm: React.FC<IGenericErrorForm> = ({ error, modalType, onSubmit, onModalClose }) => {
	const handleConfirm = () => {
		onSubmit(error);
		onModalClose(modalType);
	};

	return (
		<div id="genericErrorForm" className="genericErrorForm">
			<p>{`${error}`}</p>
			<div className="button-group justify-content-center">
				<button className="btn btn-danger" data-bs-dismiss="modal" onClick={handleConfirm}>
					OK
				</button>
			</div>
		</div>
	);
};

export default GenericErrorForm;
