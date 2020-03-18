<!DOCTYPE html>
<html lang="en">
    <head>
        <meta charset="utf-8" />
        <title>BeegoDemo后台管理</title>
        <meta name="viewport" content="width=device-width, initial-scale=1.0">
        <meta content="A fully featured admin theme which can be used to build CRM, CMS, etc." name="description" />
        <meta http-equiv="X-UA-Compatible" content="IE=edge" />
        <!-- App favicon -->

        <!-- DataTables -->
        <link href="/static/assets/libs/datatables/dataTables.bootstrap4.min.css" rel="stylesheet" type="text/css"/>
        <link href="/static/assets/libs/datatables/responsive.bootstrap4.min.css" rel="stylesheet" type="text/css"/>

        <!-- App css -->
        <link href="/static/assets/css/bootstrap.min.css" rel="stylesheet" type="text/css" />
        <link href="/static/assets/css/icons.min.css" rel="stylesheet" type="text/css" />
        <link href="/static/assets/css/app.min.css" rel="stylesheet" type="text/css" />

    </head>

    <body>

        <!-- Begin page -->
        <div id="wrapper">

            <!-- Topbar Start -->
            <div class="navbar-custom">
                <ul class="list-unstyled topnav-menu float-right mb-0">

{{range $index, $elem:= .TopMenu}}

  <li class="dropdown notification-list">
         <a class="nav-link  mr-0 waves-effect waves-light"  href="/index/{{$elem.AuthCode}}" role="button" aria-haspopup="false" aria-expanded="false">
               <span class="pro-user-name ml-1">
            {{$elem.AuthName}}
              </span>
          </a>
      </li>

{{end}}





                    <li class="dropdown notification-list">
                        <a class="nav-link dropdown-toggle nav-user mr-0 waves-effect waves-light" data-toggle="dropdown" href="#" role="button" aria-haspopup="false" aria-expanded="false">

                            <span class="pro-user-name ml-1">
                               {{.userinfo.Name}} <i class="mdi mdi-chevron-down"></i>
                            </span>
                        </a>
                        <div class="dropdown-menu dropdown-menu-right profile-dropdown ">
                            <!-- item-->
                            <div class="dropdown-item noti-title">
                                <h6 class="m-0">
                                    Hello, {{.userinfo.Name}}  !
                                </h6>
                            </div>

                            <div class="dropdown-divider"></div>

                            <!-- item-->
                            <a href="/logout" class="dropdown-item notify-item">
                                <i class="dripicons-power"></i>
                                <span>退出</span>
                            </a>

                        </div>
                    </li>

                </ul>


            </div>
            <!-- end Topbar -->

            <!-- ========== Left Sidebar Start ========== -->
            <div class="left-side-menu">

                <div class="slimscroll-menu">
                    <!--- Sidemenu -->
                    <div id="sidebar-menu">
                        <ul class="metismenu" id="side-menu">
{{range $indexleft, $leftElem:= .LeftMenu}}

  <li >
        <a href="javascript: void(0);">
                                    <i class="dripicons-view-apps"></i>
                                    <span> {{$leftElem.AuthName}} </span>
                                     <span class="menu-arrow"></span>
          </a>
          {{if $.leftElem.SubList}}
          {{else}}

               <ul class="nav-second-level" >
                    {{range $indexleft2, $leftElem2:= $leftElem.SubList}}
                        <li>
                            <a href="">{{$leftElem2.AuthName}} </a>
                        </li>
                    {{end}}
              </ul>
        {{end}}
      </li>

{{end}}
                            <li>
                                <a href="javascript: void(0);">
                                    <i class="dripicons-view-apps"></i>
                                    <span> Layouts </span>
                                    <span class="menu-arrow"></span>
                                </a>
                                <ul class="nav-second-level" aria-expanded="false">
                                    <li>
                                        <a href="layouts-dark-header.html">Dark Header</a>
                                    </li>
                                    <li>
                                        <a href="layouts-sidebar-collapsed.html">Sidebar Collapsed</a>
                                    </li>
                                    <li>
                                        <a href="layouts-dark-sidebar.html">Dark Sidebar</a>
                                    </li>
                                    <li>
                                        <a href="layouts-small-sidebar.html">Small Sidebar</a>
                                    </li>
                                    <li>
                                        <a href="layouts-boxed.html">Boxed Layout</a>
                                    </li>
                                </ul>
                            </li>

                        </ul>

                    </div>
                    <!-- End Sidebar -->

                    <div class="clearfix"></div>

                </div>
                <!-- Sidebar -left -->

            </div>
            <!-- Left Sidebar End -->

            <!-- ============================================================== -->
            <!-- Start Page Content here -->
            <!-- ============================================================== -->

            <div class="content-page">
                <div class="content">

                    <!-- Start Content-->
                    <div class="container-fluid">

                        <!-- start page title -->
                        <div class="row">
                            <div class="col-12">
                                <div class="page-title-box">
                                    <div class="page-title-right">
                                        <ol class="breadcrumb m-0">
                                            <li class="breadcrumb-item"><a href="javascript: void(0);">BeegoDemo</a></li>
                                            <li class="breadcrumb-item active">首页</li>
                                        </ol>
                                    </div>
                                    <h4 class="page-title">BeegoDemo</h4>
                                </div>
                            </div>
                        </div>
                        <!-- end page title -->


    <div class="row">
                            <div class="col-12">
                                <div class="card-box">
                                    <h4 class="header-title"> beegodemo</h4>
                                    <p class="sub-header">Welcome beegodemo</p>

                                </div> <!-- end card-box-->
                            </div> <!-- end col-->
                        </div>
                        <!-- end row-->
                    </div> <!-- container -->

                </div> <!-- content -->

                <!-- Footer Start -->
                <footer class="footer">
                    <div class="container-fluid">
                        <div class="row">
                            <div class="col-md-12 text-center">
                                2020 © Beegodemo
                            </div>
                        </div>
                    </div>
                </footer>
                <!-- end Footer -->

            </div>

            <!-- ============================================================== -->
            <!-- End Page content -->
            <!-- ============================================================== -->

        </div>
        <!-- END wrapper -->

        <!-- Vendor js -->
        <script src="/static/assets/js/vendor.min.js"></script>

        <!-- KNOB JS -->
        <script src="/static/assets/libs/jquery-knob/jquery.knob.min.js"></script>
        <!-- Chart JS -->
        <script src="/static/assets/libs/chart-js/Chart.bundle.min.js"></script>
        <!-- DatePicker JS -->
        <script src="/static/assets/libs/bootstrap-datepicker/bootstrap-datepicker.min.js"></script>

        <!-- Datatable js -->
        <script src="/static/assets/libs/datatables/jquery.dataTables.min.js"></script>
        <script src="/static/assets/libs/datatables/dataTables.bootstrap4.min.js"></script>
        <script src="/static/assets/libs/datatables/dataTables.responsive.min.js"></script>
        <script src="/static/assets/libs/datatables/responsive.bootstrap4.min.js"></script>

        <!-- Dashboard Init JS -->
        <script src="/static/assets/js/pages/dashboard.init.js"></script>

        <!-- App js -->
        <script src="/static/assets/js/app.min.js"></script>

    </body>
</html>