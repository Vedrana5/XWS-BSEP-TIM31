import { HttpErrorResponse } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormControl, FormGroup, Validators } from '@angular/forms';
import { MatDialogRef } from '@angular/material/dialog';
import { MatSnackBar } from '@angular/material/snack-bar';
import { Router } from '@angular/router';
import { Comment } from 'src/app/interfaces/comment-dto';
import { CompanyService } from 'src/app/services/company.service';

@Component({
  selector: 'app-comment',
  templateUrl: './comment.component.html',
  styleUrls: ['./comment.component.css']
})
export class CommentComponent implements OnInit {

  comment!: Comment;
  createForm!: FormGroup;
  cid!: string;
  constructor(private _formBuilder: FormBuilder,
    private companyService: CompanyService,
    private router: Router,
    public dialogRef: MatDialogRef<CommentComponent>,
    private _snackBar: MatSnackBar) {

    this.comment = {} as Comment
    this.createForm = this._formBuilder.group({
      comment: new FormControl('', Validators.required)
    })
  }

  ngOnInit(): void {
    console.log(this.router.url);
    this.cid = this.router.url.substring(13);
  }

  add() {

    this.createCommentRequest();
    console.log(this.createCommentRequest);

    this.companyService.createComment(this.comment).subscribe({
      next: (res) => {

        this.dialogRef.close({ event: "Created Job offer", data: res });
        this._snackBar.open(
          'You have created a job offer.',
          'Dismiss', {
          duration: 3000
        });
        window.location.reload();
      },
      error: (err: HttpErrorResponse) => {

        this._snackBar.open(err.error.message + "!", 'Dismiss', {
          duration: 3000
        });
      },
      complete: () => console.info('complete')
    });
  }

  createCommentRequest() {
    console.log(this.createForm.value.Name);

    this.comment.comment = this.createForm.value.comment;
    this.comment.email = localStorage.getItem("email")!;
    this.comment.companyId = parseInt(this.cid);
    console.log(this.cid)
  }

}
