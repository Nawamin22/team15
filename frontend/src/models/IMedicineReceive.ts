import { ZoneInterface } from "./IZone";
import { UserInterface } from "./IUser";
import { MedicineLabelsInterface } from "./IMedicineLabel";

export interface MedicineReceiveInterface {

    ID: number,
    MedicineReceiveNo: number,
    
    RecievedDate: Date | null,

    PharmacistID: number,
    Pharmacist: UserInterface,


    ZoneID: number,
	Zone: ZoneInterface,

    MedicineLabelID: number,
    MedicineLable: MedicineLabelsInterface,


}