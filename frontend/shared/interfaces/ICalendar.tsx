import { Dispatch, SetStateAction } from "react";
import IHabit from "./IHabit";

export default interface ICalendar {
    currentSelectedHabit: IHabit | null;
    completionDatesCounter: number;
    setCompletionDatesCompletionDatesCounter: Dispatch<SetStateAction<number>>;
    setCompletionDates: Dispatch<SetStateAction<string[]>>;
    completionDates: string[];
}
