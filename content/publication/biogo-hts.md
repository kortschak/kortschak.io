+++
abstract = "bíogo/hts provides a Go native implementation of the SAM specification (The SAM/BAM Format Specification Working Group 2016) for SAM and BAM alignment formats (Li et al. 2012) commonly used for representation of high throughput genomic data, the BAI, CSI and tabix indexing formats, and the BGZF blocked compression format. The bíogo/hts packages perform parallelized read and write operations and are able to cache recent reads according to user-specified caching methods. The parallelisation approach used by the bíogo/hts package is influenced by the approach of the D implementation, sambamba by Tarazov et al. (Tarasov et al. 2015). The bíogo/hts APIs have been constructed to provide a consistent interface to sequence alignment data and the underlying compression system in order to aid ease of use and tool development."
abstract_short = ""
authors = ["R Daniel Kortschak", "Brent S Pedersen", "David L Adelson"]
date = "2017-02-05"
image_preview = ""
math = true
publication_types = ["2"]
publication = "The Journal of Open Source Software"
publication_short = ""
featured = false
title = "bíogo/hts: high throughput sequence handling for the Go language"
url_code = "https://github.com/biogo/hts"
url_dataset = ""
url_pdf = "https://github.com/openjournals/joss-papers/raw/master/joss.00168/10.21105.joss.00168.pdf"
url_project = ""
url_slides = ""
url_video = ""

[[links]]
name = "HTML"
url = "http://joss.theoj.org/papers/10.21105/joss.00168"

+++

Tarasov A., Vilella A., Cuppen E., Nijman I. J., and Prins P. 2015. "Sambamba: Fast Processing of Ngs Alignment Formats." Bioinformatics 31 (12): 2032–4. doi:[10.1093/bioinformatics/btv098](https://doi.org/10.1093/bioinformatics/btv098).

The SAM/BAM Format Specification Working Group. 2016. "Sequence Alignment/Map Format Specification." https://samtools.github.io/hts-specs/SAMv1.pdf.

Li H., Handsaker B., Wysoker A., Fennell T., Ruan J., Homer N., Marth G., Abecasis G., Durbin R., and 1000 Genome Project Data Processing Subgroup. 2012. "The Sequence Alignment/Map Format and Samtools." Bioinformatics 25 (16): 2078–9. doi:[10.1093/bioinformatics/btp352](https://doi.org/10.1093/bioinformatics/btp352).

