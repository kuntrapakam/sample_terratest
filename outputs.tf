output "bucket_id" {
    value = aws_s3_bucket.terratest-bucket.id
}

output "bucket_arn" {
    value = aws_s3_bucket.terratest-bucket.arn
}