export interface JobOffer {
    id: string,
    position: string,
    description: string,
    dateCreated: string,
    dueDate: string | null,
    companyName: string,
    companyId: number
}