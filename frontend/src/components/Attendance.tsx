import React, { useEffect, useState } from "react";
import Container from '@mui/material/Container'
import TableCell from '@mui/material/TableCell';
import { Box, Grid, Select, TextField, Typography, Table, TableHead, TableRow, TableBody, IconButton, Alert } from '@mui/material'
import Button from '@mui/material/Button'
import { Link as RouterLink } from "react-router-dom";
import TableContainer from '@mui/material/TableContainer';
import moment from 'moment';
import DeleteIcon from '@mui/icons-material/Delete';
import Snackbar from '@mui/material/Snackbar'
import MuiAlert, { AlertProps } from '@mui/material/Alert';

import { AttendanceInterface } from "../models/IAttendance";
import { format } from "path";
function Attendances() {

    
    const [success, setSuccess] = useState(false);
    const [error, setError] = useState(false);
    const [ErrorMessage, setErrorMessage] = React.useState("");
    const [attendance, SetAttendance] = React.useState<AttendanceInterface[]>([]);

    const getAttendance = async () => {
        const apiUrl = "http://localhost:8080/attendances";
        const requestOptions = {
            method: "GET",
            headers: { "Content-Type": "application/json" },
        };
        fetch(apiUrl, requestOptions)
            .then((response) => response.json())
            .then((res) => {
                console.log(res.data);
                if (res.data) {
                    SetAttendance(res.data);
                }
            });
    };
    const DeleteAttendance = async (id: string | number | undefined) => {
        const apiUrl = "http://localhost:8080";
        const requestOptions = {
            method: "DELETE",
            headers: {
                Authorization: `Bearer ${localStorage.getItem("token")}`,
                "Content-Type": "application/json",
            },
        };

        fetch(`${apiUrl}/attendance/${id}`, requestOptions)
            .then((response) => response.json())
            .then(
                (res) => {
                    if (res.data) {
                        setSuccess(true)
                        console.log("ยกเลิกสำเร็จ")
                        setErrorMessage("")
                    }
                    else {
                        setErrorMessage(res.error)
                        setError(true)
                        console.log("ยกเลิกไม่สำเร็จ")
                    }
                    getAttendance();
                }
            )
    }


    useEffect(() => {
        getAttendance();
    }, []);
    
    const Alert = React.forwardRef<HTMLDivElement, AlertProps>(function Alert(
        props,
        ref,
    ) {
        return <MuiAlert elevation={6} ref={ref} variant="filled" {...props} />;
    });

    const handleClose = (res: any) => {
        if (res === "clickaway") {
            return;
        }
        setSuccess(false);
        setError(false);
    };

    return (

        <div>

            <Container maxWidth="md">
            <Snackbar open={success} autoHideDuration={6000} onClose={handleClose}>
                    <Alert onClose={handleClose} severity="success">
                        บันทึกข้อมูลสำเร็จ
                    </Alert>
                </Snackbar>
                <Snackbar open={error} autoHideDuration={6000} onClose={handleClose}>
                    <Alert onClose={handleClose} severity="error">
                        บันทึกข้อมูลไม่สำเร็จ: {ErrorMessage}
                    </Alert>
                </Snackbar>

                <Box

                    display="flex"

                    sx={{

                        marginTop: 2,

                    }}

                >

                    <Box flexGrow={1}>

                        <Typography
                            component="h2"
                            variant="h6"
                            color="primary"
                            gutterBottom
                        >
                            ข้อมูลการสั่งซื้อ
                        </Typography>
                    </Box>
                    <Box>

                        <Button

                            component={RouterLink}

                            to="/attendance"

                            variant="contained"

                            color="primary"

                        >

                            เพิ่มการสั่งซื้อ

                        </Button>

                    </Box>

                </Box>

                <TableContainer >

                    <Table aria-label="simple table">

                        <TableHead>
                            <TableRow>
                                <TableCell align="center" width="5%">
                                    ลำดับ
                                </TableCell>
                                <TableCell align="center" width="15%">
                                    หน้าที่
                                </TableCell>
                                <TableCell align="center" width="5%">
                                    เบอร์โทร
                                </TableCell>
                                
                                <TableCell align="center" width="10%">
                                    หมายเหตุ
                                </TableCell>
                                <TableCell align="center" width="20%">
                                    ช่วงเวลาเข้าเวร
                                </TableCell>
                                <TableCell align="center" width="15%">
                                    ผู้บันทึก
                                </TableCell>
                                <TableCell align="center" width="15%">
                                    วันที่บันทึก
                                </TableCell>
                            </TableRow>

                        </TableHead>

                        <TableBody>
                            {attendance.map((attendance: AttendanceInterface) => (
                                <TableRow key={attendance.ID}>
                                    <TableCell align="center" > {attendance.ID}            </TableCell>
                                    <TableCell align="center" > {attendance.Stat.Name}    </TableCell>
                                    <TableCell align="center" > {attendance.Phone}    </TableCell>
                                    <TableCell align="center" > {attendance.Description}         </TableCell>
                                    <TableCell align="center" > {attendance.Shift.Name}     </TableCell>
                                    <TableCell align="center" > {attendance.Pharmacist.Name}     </TableCell>                                    
                                    <TableCell align="center" > {moment(attendance.Datetime).format('DD MMMM yyyy')}     </TableCell>
                                    <TableCell align="center">
                                        <IconButton aria-label="delete" vertical-align="middle" onClick={() => DeleteAttendance(attendance.ID)}><DeleteIcon /></IconButton >
                                    </TableCell>

                                    <TableCell align="center">
                                        <Button
                                            variant='outlined'
                                            color="primary"
                                            component={RouterLink}
                                            to={"/attendanceUpdate/" + attendance.ID}
                                        >
                                            แก้ไขข้อมูล
                                        </Button>
                                    </TableCell>
                                
                                
                                
                                
                                
                                
                                
                                
                                
                                </TableRow>
                            ))}
                        </TableBody>

                    </Table>

                </TableContainer>

            </Container>

        </div>
    );
}



export default Attendances;