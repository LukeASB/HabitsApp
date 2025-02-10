export const FormatDate = (date: string = "") => {
    if (!date) return date;

    return date.split("T")[0].split("-").reverse().join("/");
};