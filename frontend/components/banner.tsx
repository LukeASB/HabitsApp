import { ReactNode } from "react";

const subtitleStyle = {
  fontSize: "x-large",
};

interface IBanner {
  children: ReactNode
}

const Banner: React.FC<IBanner> = ({ children }: { children: ReactNode}) => {
  return (
    <header className="row mb-4">
      <div className="col-12 mt-5" style={subtitleStyle}>
        {children}
      </div>
    </header>
  );
};

export default Banner;
